package common

import (
	"context"
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/pb/common"
)

type Common struct {
	svc *service.Service
}

func NewCommon(svc *service.Service) *Common {
	return &Common{svc}
}

func (c *Common) GetStatus(ctx *gin.Context) {
	res, _ := c.svc.GetStatus(context.Background(), &pbCommon.EmptyRequest{})

	ctx.JSON(http.StatusOK, gin.H{
		"enabled": res.Enabled,
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
