package common

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func CommonRoutes(router *gin.Engine) {
	h := handler.NewHandler()

	commonGroup := router.Group("/common")
	{
		commonGroup.GET("/on", h.Common.On)
		commonGroup.GET("/off", h.Common.Off)
	}
}
