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
	status, err := c.svc.GetStatus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"enabled": status,
	})
}

func (c *Common) Enable(ctx *gin.Context) {
	err := c.svc.Enable()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Common) Disable(ctx *gin.Context) {
	err := c.svc.Disable()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
