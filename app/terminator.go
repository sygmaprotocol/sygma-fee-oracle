// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package app

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (a *FeeOracleApp) terminate() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGSTOP, syscall.SIGTERM, syscall.SIGQUIT) //nolint

	go func() {
		defer a.appTerminationChecker.Done()
		for sig := range sigs {
			a.log.Warnf("terminate fee oracle due to %s", sig.String())
			a.preTerminationCheck()
			break
		}
		a.log.Info("bye-bye")
	}()
}

func (a *FeeOracleApp) preTerminationCheck() {
	a.shutdownHttpServer()

	a.stopCronJobs()

	a.stopStore()

	a.log.Info("pre-termination checking is done, preparing final termination...")
	for i := a.base.GetConfig().FinishUpTimeConfig(); i >= 1; i-- { // give some time here to allow any internal jobs finish
		if i <= 10 {
			a.log.Warn("shutdown in ", i)
		}
		time.Sleep(1 * time.Second)
	}
}
