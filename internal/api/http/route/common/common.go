package common

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseCommonRoutes(r *gin.Engine, h *handler.Handler) {
	commonGroup := r.Group("/common")

	commonGroup.GET("/on", h.Common.On)
	commonGroup.GET("/off", h.Common.Off)
}
