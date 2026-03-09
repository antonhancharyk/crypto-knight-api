package position

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	entity "github.com/antongoncharik/crypto-knight-api/internal/entity/position"
	"github.com/gin-gonic/gin"
)

type PositionService interface {
	GetPositions() (entity.Positions, error)
}

type Position struct {
	svc PositionService
}

func New(svc PositionService) *Position {
	return &Position{svc: svc}
}

// GetAll godoc
// @Summary      Get positions
// @Tags         position
// @Produce      json
// @Success      200  {object}  entity.Positions
// @Failure      500  {object}  response.ErrorResponse
// @Router       /position [get]
func (r *Position) GetAll(ctx *gin.Context) {
	res, err := r.svc.GetPositions()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.OK(ctx, res)
}
