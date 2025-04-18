package klines

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	klines := req.Group("/klines")

	klines.GET("", hdl.Klines.Get)
}
