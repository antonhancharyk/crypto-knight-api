package tracks

import (
	"fmt"
	"strings"
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/jmoiron/sqlx"
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
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, created_at, is_order, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices from tracks where created_at between $1 and $2 and symbol = $3 order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol)

		return tracksData, err
	}

	if queryParams.Full && queryParams.Symbol == "" {
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, created_at, is_order, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices from tracks where created_at between $1 and $2 order by created_at desc", queryParams.From, queryParams.To)

		return tracksData, err
	}

	if queryParams.Symbol != "" {
		err := t.db.Select(&tracksData, "with ranked_tracks as (select id, symbol, high_price, low_price, created_at, lag(high_price) over (partition by symbol order by created_at) as prev_high_price, lag(low_price) over (partition by symbol order by created_at) as prev_low_price, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices from tracks where created_at between $1 and $2 and symbol = $3) select symbol, high_price, low_price, created_at, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices from ranked_tracks where high_price != prev_high_price or low_price != prev_low_price or prev_high_price is null order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol)

		return tracksData, err
	}

	err := t.db.Select(&tracksData, "with ranked_tracks as (select id, symbol, high_price, low_price, created_at, lag(high_price) over (partition by symbol order by created_at) as prev_high_price, lag(low_price) over (partition by symbol order by created_at) as prev_low_price, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices from tracks where created_at between $1 and $2) select symbol, high_price, low_price, created_at, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices from ranked_tracks where high_price != prev_high_price or low_price != prev_low_price or prev_high_price is null order by created_at desc", queryParams.From, queryParams.To)

	return tracksData, err
}

func (t *Tracks) Create(track track.Track) error {
	var err error
	if (track.CreatedAt == time.Time{}) {
		_, err = t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price, is_order, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, track.Symbol, track.HighPrice, track.LowPrice, track.IsOrder, track.HighCreatedAt, track.LowCreatedAt, track.HighPrices, track.LowPrices, track.TakeProfitHighPrices, track.TakeProfitLowPrices)
	} else {
		_, err = t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price, created_at, is_order, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, track.Symbol, track.HighPrice, track.LowPrice, track.CreatedAt, track.IsOrder, track.HighCreatedAt, track.LowCreatedAt, track.HighPrices, track.LowPrices, track.TakeProfitHighPrices, track.TakeProfitLowPrices)
	}

	return err
}

func (t *Tracks) CreateBulk(tracks []track.Track) error {
	var placeholders []string
	var values []interface{}
	for i, track := range tracks {
		if (track.CreatedAt == time.Time{}) {
			placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*10+1, i*10+2, i*10+3, i*10+4, i*10+5, i*10+6, i*10+7, i*10+8, i*10+9, i*10+10))
			values = append(values, track.Symbol, track.HighPrice, track.LowPrice, track.IsOrder, track.HighCreatedAt, track.LowCreatedAt, track.HighPrices, track.LowPrices, track.TakeProfitHighPrices, track.TakeProfitLowPrices)
		} else {
			placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*11+1, i*11+2, i*11+3, i*11+4, i*11+5, i*11+6, i*11+7, i*11+8, i*11+9, i*11+10, i*11+11))
			values = append(values, track.Symbol, track.HighPrice, track.LowPrice, track.IsOrder, track.HighCreatedAt, track.LowCreatedAt, track.HighPrices, track.LowPrices, track.TakeProfitHighPrices, track.TakeProfitLowPrices, track.CreatedAt)
		}
	}

	var err error
	if (tracks[0].CreatedAt == time.Time{}) {
		query := fmt.Sprintf(`
		INSERT INTO tracks (symbol, high_price, low_price, is_order, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices)
		VALUES %s
	`, strings.Join(placeholders, ","))

		_, err = t.db.Exec(query, values...)
	} else {
		query := fmt.Sprintf(`
		INSERT INTO tracks (symbol, high_price, low_price, is_order, high_created_at, low_created_at, high_prices, low_prices, take_profit_high_prices, take_profit_low_prices, created_at)
		VALUES %s
	`, strings.Join(placeholders, ","))

		_, err = t.db.Exec(query, values...)
	}

	return err
}
