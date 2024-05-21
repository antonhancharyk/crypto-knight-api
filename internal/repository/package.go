package repository

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository/common"
	"github.com/antongoncharik/crypto-knight-api/internal/repository/tracks"
	"github.com/jmoiron/sqlx"
)

type Common interface {
	On() error
	Off() error
	GetStatus() (bool, error)
}

type Tracks interface {
	GetAll(queryParams track.QueryParams) ([]track.Track, error)
	Create(track track.Track) error
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
