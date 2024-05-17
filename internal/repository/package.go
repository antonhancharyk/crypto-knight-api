package repository

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository/common"
	"github.com/antongoncharik/crypto-knight-api/internal/repository/tracks"
	"github.com/jmoiron/sqlx"
)

type Common interface {
	On()
	Off()
	GetStatus() bool
}

type Tracks interface {
	GetAll() []track.Track
	Create(track track.Track)
}

type Repository struct {
	Common
	Tracks
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Common: common.New(db),
		Tracks: tracks.New(db),
	}
}
