package klines

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/kline"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Kline struct {
	svc *service.Service
}

func New(svc *service.Service) *Kline {
	return &Kline{svc}
}

func (r *Kline) Get(ctx *gin.Context) {
	var queryParams kline.QueryParams

	err := ctx.ShouldBindQuery(&queryParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := r.svc.Kline.Get(queryParams.Symbol)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
