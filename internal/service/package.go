package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	authSvc "github.com/antongoncharik/crypto-knight-api/internal/service/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
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
}

type Auth interface {
	ValidateToken(token string) error
}

type Service struct {
	Common
	Tracks
	Auth
}

func New(repo *repository.Repository, keys auth.RSAKeys) *Service {
	return &Service{
		Common: common.New(repo),
		Tracks: tracks.New(repo),
		Auth:   authSvc.New(repo, keys),
	}
}
