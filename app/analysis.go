// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package app

import (
	"github.com/ChainSafe/sygma-fee-oracle/config"
	"net/http"
)
import _ "net/http/pprof" // #nosec

func (a *FeeOracleApp) goroutineMemoryLeakChecker() {
	if a.base.GetEnv() != config.AppEvmProd {
		go func() {
			a.base.GetLogger().Info(http.ListenAndServe("localhost:6060", nil))
		}()
	}
}
