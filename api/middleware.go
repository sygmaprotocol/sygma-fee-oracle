package api

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterSetup(ginInstance *gin.Engine, appBase *base.FeeOracleAppBase) {
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
