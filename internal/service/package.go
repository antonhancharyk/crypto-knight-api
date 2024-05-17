package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service/tracks"
)

type Common interface {
	Enable()
	Disable()
	GetStatus() bool
}

type Tracks interface {
	GetAll() []track.Track
	Create(track track.Track)
}

type Service struct {
	Common
	Tracks
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Common: common.New(repo),
		Tracks: tracks.New(repo),
	}
}
