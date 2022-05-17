// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// AddRouterPathsRoot registers root api routers
func AddRouterPathsRoot(rootRouter *gin.RouterGroup, apiHandler *Handler, log *logrus.Entry) {
	apiHandler.log = log.WithField("api", "root")

	rootRouter.GET("/health", apiHandler.healthCheck)
}

func (h *Handler) healthCheck(c *gin.Context) {
	ginSuccessReturn(c, http.StatusOK, "OK")
}
