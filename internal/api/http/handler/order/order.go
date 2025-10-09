package order

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Order struct {
	svc *service.Service
}

func New(svc *service.Service) *Order {
	return &Order{svc}
}

func (r *Order) GetAll(ctx *gin.Context) {
	res, err := r.svc.Order.GetOpenOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
