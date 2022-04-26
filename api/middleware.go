// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/ChainSafe/chainbridge-fee-oracle/consensus"
	"github.com/ChainSafe/chainbridge-fee-oracle/identity"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterSetup(ginInstance *gin.Engine, identity *identity.OracleIdentityOperator, consensus *consensus.Consensus,
	gasPriceStore *store.GasPriceStore, conversionRateStore *store.ConversionRateStore, conf *config.Config, log *logrus.Entry) {
	apiHandlerV1 := &Handler{
		conf:                conf,
		identity:            identity,
		consensus:           consensus,
		gasPriceStore:       gasPriceStore,
		conversionRateStore: conversionRateStore,
	}

	// global scope middleware
	ginInstance.Use(gin.Recovery(), corsSetup)

	// router path register
	v1 := ginInstance.Group("/v1")
	v1Rate := v1.Group("/rate")

	// v1 apis groups middleware
	v1Rate.Use()

	// v1 router group map
	v1routers := make(map[string]*gin.RouterGroup)
	v1routers["rate"] = v1Rate

	AddRouterPathsV1(v1routers, apiHandlerV1, log)
}

/* middlewares */

func corsSetup(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Language, Accept-Encoding, X-CSRF-Token, X-Auth-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, PATCH, POST, GET, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
