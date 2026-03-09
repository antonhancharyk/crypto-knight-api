package auth

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	ValidateToken(token string) error
}

type Auth struct {
	svc AuthService
}

func New(svc AuthService) *Auth {
	return &Auth{svc: svc}
}

// ValidateToken godoc
// @Summary      Validate auth token
// @Tags         auth
// @Param        token  query  string  true  "JWT token"
// @Success      200
// @Failure      400  {object}  response.ErrorResponse
// @Failure      403  {object}  response.ErrorResponse
// @Router       /auth/validate [get]
func (a *Auth) ValidateToken(ctx *gin.Context) {
	var paramsToken auth.QueryParams

	if err := ctx.ShouldBindQuery(&paramsToken); err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	if err := a.svc.ValidateToken(paramsToken.Token); err != nil {
		response.Error(ctx, http.StatusForbidden, err)
		return
	}

	ctx.Status(http.StatusOK)
}
