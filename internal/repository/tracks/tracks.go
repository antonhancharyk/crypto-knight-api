package tracks

import (
	"fmt"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/jmoiron/sqlx"
)

type Tracks struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Tracks {
	return &Tracks{db: db}
}

func (t *Tracks) GetAll(queryParams track.QueryParams) ([]track.Track, error) {
	var tracksData []track.Track

	err := t.db.Select(&tracksData, "select symbol, high_price, low_price, created_at from tracks where created_at between $1 AND $2 order by created_at desc", queryParams.From, queryParams.To)
	fmt.Println("db", tracksData)
	return tracksData, err
}

func (t *Tracks) Create(track track.Track) error {
	_, err := t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price)
	VALUES ($1, $2, $3)`, track.Symbol, track.HighPrice, track.LowPrice)

	return err
}
