package balance

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Balance struct {
	svc *service.Service
}

func New(svc *service.Service) *Balance {
	return &Balance{svc}
}

func (r *Balance) Get(ctx *gin.Context) {
	res, err := r.svc.Balance.Get()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
