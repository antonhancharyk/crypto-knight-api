package common

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Common struct {
	svc *service.Service
}

func New(svc *service.Service) *Common {
	return &Common{svc}
}

func (c *Common) GetStatus(ctx *gin.Context) {
	status := c.svc.GetStatus()

	ctx.JSON(http.StatusOK, gin.H{
		"enabled": status,
	})
}

func (c *Common) Enable(ctx *gin.Context) {
	c.svc.Enable()
	ctx.Status(http.StatusOK)
}

func (c *Common) Disable(ctx *gin.Context) {
	c.svc.Disable()
	ctx.Status(http.StatusOK)
}
