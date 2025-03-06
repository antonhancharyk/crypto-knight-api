package position

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Position struct {
	svc *service.Service
}

func New(svc *service.Service) *Position {
	return &Position{svc}
}

func (r *Position) GetAll(ctx *gin.Context) {
	res, err := r.svc.Position.GetPositions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
