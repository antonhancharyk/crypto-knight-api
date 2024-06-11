package auth

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	svc *service.Service
}

func New(svc *service.Service) *Auth {
	return &Auth{svc}
}

func (a *Auth) ValidateToken(ctx *gin.Context) {
	var paramsToken auth.QueryParams

	err := ctx.ShouldBindQuery(&paramsToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.svc.ValidateToken(paramsToken.Token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
