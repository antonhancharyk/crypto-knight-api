package tracks

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
)

type Tracks struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Tracks {
	return &Tracks{repo: repo}
}

func (t *Tracks) GetAll() []track.Track {
	return t.repo.Tracks.GetAll()
}

func (t *Tracks) Create(track track.Track) {
	t.repo.Tracks.Create(track)
}
