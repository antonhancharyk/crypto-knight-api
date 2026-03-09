package order

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/order"
	"github.com/gin-gonic/gin"
)

type OrderService interface {
	GetOpenOrders() ([]order.Order, error)
}

type Order struct {
	svc OrderService
}

func New(svc OrderService) *Order {
	return &Order{svc: svc}
}

// GetAll godoc
// @Summary      Get open orders
// @Tags         orders
// @Produce      json
// @Success      200  {array}   order.Order
// @Failure      500  {object}  response.ErrorResponse
// @Router       /orders [get]
func (r *Order) GetAll(ctx *gin.Context) {
	res, err := r.svc.GetOpenOrders()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.OK(ctx, res)
}
