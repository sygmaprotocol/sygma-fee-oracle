// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package cronjob

import (
	"github.com/ChainSafe/sygma-fee-oracle/base"
	"github.com/ChainSafe/sygma-fee-oracle/oracle"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Cronjob struct {
	base    *base.FeeOracleAppBase
	cronJob *cron.Cron

	conversionRateOracles map[string]*oracle.ConversionRateOracleOperator
	gasPriceOracles       map[string]*oracle.GasPriceOracleOperator

	conversionRateStore *store.ConversionRateStore
	gasPriceStore       *store.GasPriceStore

	log *logrus.Entry
}

func NewCronJobs(base *base.FeeOracleAppBase, conversionRateOracles map[string]*oracle.ConversionRateOracleOperator,
	gasPriceOracles map[string]*oracle.GasPriceOracleOperator, conversionRateStore *store.ConversionRateStore,
	gasPriceStore *store.GasPriceStore, logger *logrus.Entry) *Cronjob {
	return &Cronjob{
		base:    base,
		cronJob: cron.New(),

		conversionRateOracles: conversionRateOracles,
		gasPriceOracles:       gasPriceOracles,

		conversionRateStore: conversionRateStore,
		gasPriceStore:       gasPriceStore,

		log: logger.WithField("cronjob", "base"),
	}
}

// CronJobInit prepares or checks anything necessary pre cron jobs started
func (c *Cronjob) CronJobInit() {
	c.log.Info("cronjob init success")
}

func (c *Cronjob) AddJob(job ScheduledJob) {
	if job.IsEnabled() {
		_, err := c.cronJob.AddFunc(job.GetJobSchedule(), job.GetOperationFunc())
		if err != nil {
			c.log.Fatalf("failed to add cron job: %s during app starting", job.JobName())
		}
	}
}

func (c *Cronjob) Start() {
	c.cronJob.Start()

	c.log.Info("internal cronjob started")
}

func (c *Cronjob) Stop() {
	c.cronJob.Stop()
}
