package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	entityBalance "github.com/antongoncharik/crypto-knight-api/internal/entity/balance"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	entity "github.com/antongoncharik/crypto-knight-api/internal/entity/position"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	authSvc "github.com/antongoncharik/crypto-knight-api/internal/service/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/service/balance"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service/entries"
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

type Balance interface {
	Get() (entityBalance.Balance, error)
}

type Service struct {
	Common
	Tracks
	Auth
	Entries
	Position
	Balance
}

func New(repo *repository.Repository, keys auth.RSAKeys, grpcClients *grpc.GRPCClients, apiClient *api.HTTPClient) *Service {
	return &Service{
		Common:   common.New(repo, grpcClients),
		Tracks:   tracks.New(repo),
		Auth:     authSvc.New(repo, keys),
		Entries:  entries.New(repo),
		Position: position.New(apiClient),
		Balance:  balance.New(apiClient)}
}
