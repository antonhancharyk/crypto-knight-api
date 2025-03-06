package auth

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	auth := req.Group("/auth")

	auth.GET("/validate", hdl.Auth.ValidateToken)
}
