package common

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/gin-gonic/gin"
)

type CommonService interface {
	GetStatus() (bool, error)
	Enable() error
	Disable() error
}

type Common struct {
	svc CommonService
}

func New(svc CommonService) *Common {
	return &Common{svc: svc}
}

type StatusResponse struct {
	Enabled bool `json:"enabled"`
}

// GetStatus godoc
// @Summary      Get bot status
// @Tags         common
// @Produce      json
// @Success      200  {object}  StatusResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /common/status [get]
func (c *Common) GetStatus(ctx *gin.Context) {
	status, err := c.svc.GetStatus()
	if err != nil {
		response.WriteError(ctx, err)
		return
	}

	response.JSON(ctx, http.StatusOK, StatusResponse{Enabled: status})
}

// Enable godoc
// @Summary      Enable bot
// @Tags         common
// @Success      200
// @Failure      500  {object}  response.ErrorResponse
// @Router       /common/on [get]
func (c *Common) Enable(ctx *gin.Context) {
	if err := c.svc.Enable(); err != nil {
		response.WriteError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// Disable godoc
// @Summary      Disable bot
// @Tags         common
// @Success      200
// @Failure      500  {object}  response.ErrorResponse
// @Router       /common/off [get]
func (c *Common) Disable(ctx *gin.Context) {
	if err := c.svc.Disable(); err != nil {
		response.WriteError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
