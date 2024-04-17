package common

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Common struct {
	svc *service.Service
}

func NewCommon(svc *service.Service) *Common {
	return &Common{svc}
}

func (c *Common) GetStatus(ctx *gin.Context) {
	status := c.svc.GetStatus()

	ctx.JSON(http.StatusOK, gin.H{
		"enabled": status,
	})
}

func (c *Common) On(ctx *gin.Context) {
	c.svc.On()
	ctx.Status(http.StatusOK)
}

func (c *Common) Off(ctx *gin.Context) {
	c.svc.Off()
	ctx.Status(http.StatusOK)
}
