package balance

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/balance"
	"github.com/gin-gonic/gin"
)

type BalanceService interface {
	Get() (balance.Balance, error)
}

type Balance struct {
	svc BalanceService
}

func New(svc BalanceService) *Balance {
	return &Balance{svc: svc}
}

// Get godoc
// @Summary      Get USDT balance
// @Tags         balance
// @Produce      json
// @Success      200  {object}  balance.Balance
// @Failure      500  {object}  response.ErrorResponse
// @Router       /balance [get]
func (r *Balance) Get(ctx *gin.Context) {
	res, err := r.svc.Get()
	if err != nil {
		response.WriteError(ctx, err)
		return
	}

	response.OK(ctx, res)
}
