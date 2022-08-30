// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

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
