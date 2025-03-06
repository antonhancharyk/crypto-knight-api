package common

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	common := req.Group("/common")

	common.GET("/status", hdl.Common.GetStatus)
	common.GET("/on", hdl.Common.Enable)
	common.GET("/off", hdl.Common.Disable)
}
