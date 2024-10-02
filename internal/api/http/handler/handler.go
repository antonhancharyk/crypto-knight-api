package handler

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/common"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/tracks"
	"github.com/antongoncharik/crypto-knight-api/internal/cache"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Common interface {
	GetStatus(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
}

type Tracks interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type Auth interface {
	ValidateToken(c *gin.Context)
}

type Handler struct {
	Common
	Tracks
	Auth
}

func New(svc *service.Service, cacheClient *cache.Cache) *Handler {
	return &Handler{
		Common: common.New(svc),
		Tracks: tracks.New(svc, cacheClient),
		Auth:   auth.New(svc),
	}
}
