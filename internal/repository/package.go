package repository

import (
	"github.com/antongoncharik/crypto-knight-api/internal/repository/common"
	"github.com/jmoiron/sqlx"
)

type Common interface {
	On()
	Off()
	GetStatus() bool
}

type Repository struct {
	Common
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Common: common.NewCommon(db),
	}
}
