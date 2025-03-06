package handler

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/common"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/entries"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/position"
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
	CreateBulk(c *gin.Context)
}

type Auth interface {
	ValidateToken(c *gin.Context)
}

type Entries interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type Position interface {
	GetAll(c *gin.Context)
}

type Handler struct {
	Common
	Tracks
	Auth
	Entries
	Position
}

func New(svc *service.Service, cacheClient *cache.Cache) *Handler {
	return &Handler{
		Common:   common.New(svc),
		Tracks:   tracks.New(svc, cacheClient),
		Auth:     auth.New(svc),
		Entries:  entries.New(svc),
		Position: position.New(svc),
	}
}
