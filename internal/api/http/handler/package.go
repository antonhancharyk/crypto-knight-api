package handler

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Common interface {
	On(c *gin.Context)
	Off(c *gin.Context)
}

type Handler struct {
	Common
}

func NewHandler(ser *service.Service) *Handler {
	return &Handler{
		Common: common.NewCommon(ser),
	}
}
