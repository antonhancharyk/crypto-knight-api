package common

import (
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Common struct {
	ser *service.Service
}

func NewCommon(ser *service.Service) *Common {
	return &Common{ser}
}

func (c *Common) On(ctx *gin.Context) {
	c.ser.On()
	ctx.Status(200)
}

func (c *Common) Off(ctx *gin.Context) {
	c.ser.Off()
	ctx.Status(200)
}
