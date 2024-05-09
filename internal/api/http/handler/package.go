package handler

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Common interface {
	GetStatus(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
}

type Handler struct {
	Common
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		Common: common.NewCommon(svc),
	}
}
