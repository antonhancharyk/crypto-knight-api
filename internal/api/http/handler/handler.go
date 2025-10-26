package handler

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/balance"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/common"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/entries"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/klines"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/order"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/position"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler/tracks"
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
	GetAllHistory(c *gin.Context)
	CreateBulkHistory(c *gin.Context)
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

type Order interface {
	GetAll(c *gin.Context)
}

type Balance interface {
	Get(c *gin.Context)
}

type Klines interface {
	Get(c *gin.Context)
}

type Handler struct {
	Common
	Tracks
	Auth
	Entries
	Position
	Balance
	Order
	Klines
}

func New(svc *service.Service) *Handler {
	return &Handler{
		Common:   common.New(svc),
		Tracks:   tracks.New(svc),
		Auth:     auth.New(svc),
		Entries:  entries.New(svc),
		Position: position.New(svc),
		Balance:  balance.New(svc),
		Order:    order.New(svc),
		Klines:   klines.New(svc),
	}
}
