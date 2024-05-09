package common

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseCommonRoutes(req *gin.Engine, hdl *handler.Handler) {
	commonGroup := req.Group("/common")

	commonGroup.GET("/status", hdl.Common.GetStatus)
	commonGroup.GET("/on", hdl.Common.Enable)
	commonGroup.GET("/off", hdl.Common.Disable)
}
