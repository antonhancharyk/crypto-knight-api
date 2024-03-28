package handler

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/common"
	"github.com/gin-gonic/gin"
)

type Common interface {
	On(c *gin.Context)
	Off(c *gin.Context)
}

type Handler struct {
	Common
}

func NewHandler() *Handler {
	return &Handler{
		Common: common.NewCommon(),
	}
}
