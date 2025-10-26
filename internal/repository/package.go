package repository

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository/common"
	"github.com/antongoncharik/crypto-knight-api/internal/repository/entries"
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
	CreateBulk(tracks []track.Track) error
	GetAllHistory(queryParams track.QueryParams) ([]track.Track, error)
	CreateBulkHistory(tracks []track.Track) error
}

type Entries interface {
	GetAll() ([]entry.Entry, error)
	Create(entry entry.Entry) error
}

type Repository struct {
	Common
	Tracks
	Entries
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Common:  common.New(db),
		Tracks:  tracks.New(db),
		Entries: entries.New(db),
	}
}
