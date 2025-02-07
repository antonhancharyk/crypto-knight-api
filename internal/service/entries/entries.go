package entries

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
)

type Entries struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Entries {
	return &Entries{repo: repo}
}

func (e *Entries) GetAll() ([]entry.Entry, error) {
	return e.repo.Entries.GetAll()
}

func (e *Entries) Create(entry entry.Entry) error {
	return e.repo.Entries.Create(entry)
}
