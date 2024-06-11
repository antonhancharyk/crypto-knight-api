package auth

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	commonGroup := req.Group("/auth")

	commonGroup.GET("/validate", hdl.Auth.ValidateToken)
}
