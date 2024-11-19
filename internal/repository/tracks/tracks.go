package tracks

import (
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Tracks struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Tracks {
	return &Tracks{db: db}
}

func (t *Tracks) GetAll(queryParams track.QueryParams) ([]track.Track, error) {
	tracksData := []track.Track{}

	if queryParams.Full && queryParams.Symbol != "" {
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, high_price_1, low_price_1, high_price_2, low_price_2, high_price_3, low_price_3, COALESCE(causes, '{}') as causes, created_at, is_order from tracks where created_at between $1 and $2 and symbol = $3 order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol)

		return tracksData, err
	}

	if queryParams.Full && queryParams.Symbol == "" {
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, high_price_1, low_price_1, high_price_2, low_price_2, high_price_3, low_price_3, COALESCE(causes, '{}') as causes, created_at, is_order from tracks where created_at between $1 and $2 order by created_at desc", queryParams.From, queryParams.To)

		return tracksData, err
	}

	if queryParams.Symbol != "" {
		err := t.db.Select(&tracksData, "with unique_tracks as (select distinct on (symbol, high_price, low_price) symbol, high_price, low_price, high_price_1, low_price_1, high_price_2, low_price_2, high_price_3, low_price_3, COALESCE(causes, '{}') as causes, created_at, is_order from tracks where created_at between $1 and $2 and symbol = $3 order by symbol, high_price, low_price, created_at asc) select * from unique_tracks order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol)

		return tracksData, err
	}

	err := t.db.Select(&tracksData, "with unique_tracks as (select distinct on (symbol, high_price, low_price) symbol, high_price, low_price, high_price_1, low_price_1, high_price_2, low_price_2, high_price_3, low_price_3, COALESCE(causes, '{}') as causes, created_at, is_order from tracks where created_at between $1 and $2 order by symbol, high_price, low_price, created_at asc) select * from unique_tracks order by created_at desc", queryParams.From, queryParams.To)

	return tracksData, err
}

func (t *Tracks) Create(track track.Track) error {
	var err error
	if (track.CreatedAt == time.Time{}) {
		_, err = t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price, high_price_1, low_price_1, high_price_2, low_price_2, high_price_3, low_price_3, causes, is_order)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 11)`, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrice1, track.LowPrice1, track.HighPrice2, track.LowPrice2, track.HighPrice3, track.LowPrice3, pq.Array(track.Causes), track.IsOrder)
	} else {
		_, err = t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price, high_price_1, low_price_1, high_price_2, low_price_2, high_price_3, low_price_3, causes, created_at, is_order)
		VALUES ($1, $2, $3, $4, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrice1, track.LowPrice1, track.HighPrice2, track.LowPrice2, track.HighPrice3, track.LowPrice3, pq.Array(track.Causes), track.CreatedAt, track.IsOrder)
	}

	return err
}
