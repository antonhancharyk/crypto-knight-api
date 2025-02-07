package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	authSvc "github.com/antongoncharik/crypto-knight-api/internal/service/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service/entries"
	"github.com/antongoncharik/crypto-knight-api/internal/service/tracks"
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

type Service struct {
	Common
	Tracks
	Auth
	Entries
}

func New(repo *repository.Repository, keys auth.RSAKeys, grpcClients *grpc.GRPCClients) *Service {
	return &Service{
		Common:  common.New(repo, grpcClients),
		Tracks:  tracks.New(repo),
		Auth:    authSvc.New(repo, keys),
		Entries: entries.New(repo),
	}
}
