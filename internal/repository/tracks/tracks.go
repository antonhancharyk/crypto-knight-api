package tracks

import (
	"log"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/jmoiron/sqlx"
)

type Tracks struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Tracks {
	return &Tracks{db: db}
}

func (t *Tracks) GetAll() []track.Track {
	var tracksData []track.Track

	err := t.db.Select(&tracksData, "select symbol, high_price, low_price, created_at from tracks order by created_at desc")
	if err != nil {
		log.Fatal(err)
	}

	return tracksData
}

func (t *Tracks) Create(track track.Track) {
	_, err := t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price)
	VALUES ($1, $2, $3)`, track.Symbol, track.HighPrice, track.LowPrice)
	if err != nil {
		log.Fatal(err)
	}
}
