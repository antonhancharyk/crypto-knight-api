package klines

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/kline"
	"github.com/gin-gonic/gin"
)

type KlineService interface {
	Get(sbl string) ([][]any, error)
}

type Kline struct {
	svc KlineService
}

func New(svc KlineService) *Kline {
	return &Kline{svc: svc}
}

// Get godoc
// @Summary      Get klines for symbol
// @Tags         klines
// @Param        symbol  query  string  true  "Trading pair (e.g. BTCUSDT)"
// @Produce      json
// @Success      200  {array}   array
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /klines [get]
func (r *Kline) Get(ctx *gin.Context) {
	var queryParams kline.QueryParams

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := r.svc.Get(queryParams.Symbol)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
