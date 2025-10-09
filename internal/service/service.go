package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	entityBalance "github.com/antongoncharik/crypto-knight-api/internal/entity/balance"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/order"
	entity "github.com/antongoncharik/crypto-knight-api/internal/entity/position"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	authSvc "github.com/antongoncharik/crypto-knight-api/internal/service/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/service/balance"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service/entries"
	kline "github.com/antongoncharik/crypto-knight-api/internal/service/klines"
	orderSvc "github.com/antongoncharik/crypto-knight-api/internal/service/order"
	"github.com/antongoncharik/crypto-knight-api/internal/service/position"
	"github.com/antongoncharik/crypto-knight-api/internal/service/tracks"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
)

type Common interface {
	Enable() error
	Disable() error
	GetStatus() (bool, error)
}

type Tracks interface {
	GetAll(queryParams track.QueryParams) ([]track.Track, error)
	Create(track track.Track) error
	CreateBulk(tracks []track.Track) error
}

type Auth interface {
	ValidateToken(token string) error
}

type Entries interface {
	GetAll() ([]entry.Entry, error)
	Create(entry entry.Entry) error
}

type Position interface {
	GetPositions() (entity.Positions, error)
}

type Order interface {
	GetOpenOrders() ([]order.Order, error)
}

type Balance interface {
	Get() (entityBalance.Balance, error)
}

type Kline interface {
	Get(sbl string) ([][]any, error)
}

type Service struct {
	Common
	Tracks
	Auth
	Entries
	Position
	Balance
	Order
	Kline
}

func New(repo *repository.Repository, keys auth.RSAKeys, grpcClients *grpc.GRPCClients, apiClient *api.HTTPClient) *Service {
	return &Service{
		Common:   common.New(repo, grpcClients),
		Tracks:   tracks.New(repo),
		Auth:     authSvc.New(repo, keys),
		Entries:  entries.New(repo),
		Position: position.New(apiClient),
		Balance:  balance.New(apiClient),
		Order:    orderSvc.New(apiClient),
		Kline:    kline.New(apiClient),
	}
}
