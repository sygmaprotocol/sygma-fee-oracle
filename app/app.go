// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package app

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/consensus"
	"github.com/ChainSafe/sygma-fee-oracle/identity"

	"github.com/ChainSafe/sygma-fee-oracle/api"
	"github.com/ChainSafe/sygma-fee-oracle/base"
	"github.com/ChainSafe/sygma-fee-oracle/cronjob"
	"github.com/ChainSafe/sygma-fee-oracle/oracle"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FeeOracleApp struct {
	base *base.FeeOracleAppBase

	appMode config.AppMode

	log *logrus.Entry

	ginInstance *gin.Engine // Gin engine with http server instance

	cronJobServer *cronjob.Cronjob

	conversionRateOracles map[string]*oracle.ConversionRateOracleOperator
	gasPriceOracles       map[string]*oracle.GasPriceOracleOperator

	conversionRateStore *store.ConversionRateStore
	gasPriceStore       *store.GasPriceStore

	// identity is the oracle identity, used when signing the oracle data
	identity *identity.OracleIdentityOperator

	// consensus is used for local and group oracle data checking and verification
	consensus *consensus.Consensus

	appTerminationChecker sync.WaitGroup
}

func NewFeeOracleApp(appBase *base.FeeOracleAppBase) *FeeOracleApp {
	// init concrete oracle services
	coinMarketCap := oracle.NewCoinMarketCap(appBase.GetConfig(), appBase.GetLogger())
	etherscan := oracle.NewEtherscan(appBase.GetConfig(), appBase.GetLogger())
	polygonscan := oracle.NewPolygonscan(appBase.GetConfig(), appBase.GetLogger())
	moonscan := oracle.NewMoonscan(appBase.GetConfig(), appBase.GetLogger())

	// register concrete oracle services in operator
	coinMarketCapConversionRateOracle := oracle.NewConversionRateOracleOperator(appBase.GetLogger(), coinMarketCap)
	etherscanGasPriceOracle := oracle.NewGasPriceOracleOperator(appBase.GetLogger(), etherscan)
	polygonscanGasPriceOracle := oracle.NewGasPriceOracleOperator(appBase.GetLogger(), polygonscan)
	moonscanGasPriceOracle := oracle.NewGasPriceOracleOperator(appBase.GetLogger(), moonscan)

	conversionRateOracles := make(map[string]*oracle.ConversionRateOracleOperator)
	conversionRateOracles[coinMarketCap.Name()] = coinMarketCapConversionRateOracle

	gasPriceOracles := make(map[string]*oracle.GasPriceOracleOperator)
	gasPriceOracles[etherscan.Name()] = etherscanGasPriceOracle
	gasPriceOracles[polygonscan.Name()] = polygonscanGasPriceOracle
	gasPriceOracles[moonscan.Name()] = moonscanGasPriceOracle

	conversionRateStore := store.NewConversionRateStore(appBase.GetStore())
	gasPriceStore := store.NewGasPriceStore(appBase.GetStore())

	oracleIdentity := identity.NewOracleIdentityOperator(appBase.GetOracleIdentity())

	app := &FeeOracleApp{
		appMode:               appBase.GetConfig().AppModeConfig(),
		base:                  appBase,
		log:                   appBase.GetLogger().WithField("app", "app"),
		ginInstance:           appBase.GetConfig().PrepareHttpServer(),
		conversionRateOracles: conversionRateOracles,
		gasPriceOracles:       gasPriceOracles,
		conversionRateStore:   conversionRateStore,
		gasPriceStore:         gasPriceStore,
		identity:              oracleIdentity,
		consensus:             consensus.NewConsensus(config.GetStrategy(appBase.GetConfig().StrategyConfig()), appBase.GetLogger()),
		cronJobServer: cronjob.NewCronJobs(appBase, conversionRateOracles, gasPriceOracles, conversionRateStore,
			gasPriceStore, appBase.GetLogger()),
		appTerminationChecker: sync.WaitGroup{},
	}

	app.base.GetLogger().Infof("running in: %s", appBase.GetEnv())
	app.base.GetLogger().Infof("running mode: %s", app.appMode)
	app.base.GetLogger().Info("fee oracle app init success")

	return app
}

func (a *FeeOracleApp) Start() {
	a.startHttpServer()

	if a.appMode != config.AppModeDebug {
		a.startCronJobs()
	}

	a.goroutineMemoryLeakChecker()

	a.appTerminationChecker.Add(1)
	a.terminate()
	a.appTerminationChecker.Wait()
}

// HTTP server
func (a *FeeOracleApp) startHttpServer() {
	go func() {
		api.RouterSetup(a.ginInstance, a.identity, a.consensus, a.gasPriceStore, a.conversionRateStore, a.base.GetConfig(), a.log)

		a.base.GetLogger().Infof("http server starts on port %s ", a.base.GetConfig().HttpServerConfig().Port)
		a.base.GetLogger().Infof("http server mode: %s ", a.base.GetConfig().HttpServerConfig().Mode)
		err := a.ginInstance.Run(a.base.GetConfig().HttpServerConfig().Port)
		if err != nil {
			a.base.GetLogger().Fatal("http server start error ", err)
		}
	}()
}

func (a *FeeOracleApp) shutdownHttpServer() {
	a.log.Warn("closing HTTP server...")

	srv := &http.Server{
		Addr:    a.base.GetConfig().HttpServerConfig().Port,
		Handler: a.ginInstance,
	}

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		a.log.Error("http server shutdown err: ", err)
	}
}

func (a *FeeOracleApp) startCronJobs() {
	a.cronJobServer.CronJobInit()

	// cron job scheduled to update gas price
	gasPriceJob := cronjob.NewJob(a.cronJobServer,
		a.base.GetConfig().CronJobConfig().UpdateGasPriceJob.Name,
		a.base.GetConfig().CronJobConfig().UpdateGasPriceJob.CheckFrequency,
		a.base.GetConfig().CronJobConfig().UpdateGasPriceJob.Enable)
	gasPriceJob.SetOperationFunc(cronjob.GasPriceJobOperation(gasPriceJob))
	a.cronJobServer.AddJob(gasPriceJob)

	// cron job scheduled to update conversion rate
	conversionRateJob := cronjob.NewJob(a.cronJobServer,
		a.base.GetConfig().CronJobConfig().UpdateConversionRateJob.Name,
		a.base.GetConfig().CronJobConfig().UpdateConversionRateJob.CheckFrequency,
		a.base.GetConfig().CronJobConfig().UpdateConversionRateJob.Enable)
	conversionRateJob.SetOperationFunc(cronjob.ConversionRateJobOperation(conversionRateJob))
	a.cronJobServer.AddJob(conversionRateJob)

	a.cronJobServer.Start()
}

func (a *FeeOracleApp) stopCronJobs() {
	a.log.Warn("stopping all cron jobs...")

	a.cronJobServer.Stop()
}

func (a *FeeOracleApp) stopStore() {
	err := a.base.GetStore().Close()
	if err != nil {
		a.log.Error(err)
	}
}
