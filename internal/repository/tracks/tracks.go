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
	var values []any
	for i, track := range tracks {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		values = append(values, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrices, track.LowPrices)
	}

	query := fmt.Sprintf(`
		INSERT INTO tracks (symbol, high_price, low_price, high_prices, low_prices)
		VALUES %s
	`, strings.Join(placeholders, ","))

	_, err := t.db.Exec(query, values...)

	return err
}

func (t *Tracks) GetAllHistory(queryParams track.QueryParams) ([]track.Track, error) {
	tracksData := []track.Track{}

	if queryParams.Full && queryParams.Symbol != "" {
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, created_at, high_prices, low_prices from tracks_history where created_at between $1 and $2 and symbol = $3 and interval = $4 order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol, queryParams.Interval)

		return tracksData, err
	}

	if queryParams.Full && queryParams.Symbol == "" {
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, created_at, high_prices, low_prices from tracks_history where created_at between $1 and $2 and interval = $3 order by created_at desc", queryParams.From, queryParams.To, queryParams.Interval)

		return tracksData, err
	}

	if queryParams.Symbol != "" {
		err := t.db.Select(&tracksData, "with ranked_tracks as (select id, symbol, high_price, low_price, created_at, lag(high_price) over (partition by symbol order by created_at) as prev_high_price, lag(low_price) over (partition by symbol order by created_at) as prev_low_price, high_prices, low_prices from tracks_history where created_at between $1 and $2 and symbol = $3 and interval = $4) select symbol, high_price, low_price, created_at, high_prices, low_prices from ranked_tracks where high_price != prev_high_price or low_price != prev_low_price or prev_high_price is null order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol, queryParams.Interval)

		return tracksData, err
	}

	err := t.db.Select(&tracksData, "with ranked_tracks as (select id, symbol, high_price, low_price, created_at, lag(high_price) over (partition by symbol order by created_at) as prev_high_price, lag(low_price) over (partition by symbol order by created_at) as prev_low_price, high_prices, low_prices from tracks_history where created_at between $1 and $2 and interval = $3) select symbol, high_price, low_price, created_at, high_prices, low_prices from ranked_tracks where high_price != prev_high_price or low_price != prev_low_price or prev_high_price is null order by created_at desc", queryParams.From, queryParams.To, queryParams.Interval)

	return tracksData, err
}

func (t *Tracks) CreateBulkHistory(tracks []track.Track) error {
	var placeholders []string
	var values []any
	for i, track := range tracks {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7))
		values = append(values, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrices, track.LowPrices, track.CreatedAt, track.Interval)
	}

	query := fmt.Sprintf(`
		INSERT INTO tracks_history (symbol, high_price, low_price, high_prices, low_prices, created_at, interval)
		VALUES %s
	`, strings.Join(placeholders, ","))

	_, err := t.db.Exec(query, values...)

	return err
}

func (t *Tracks) GetLastTracks() ([]track.Track, error) {
	res := []track.Track{}
	err := t.db.Select(&res, "select distinct on (symbol) symbol, high_price, low_price, high_prices, low_prices, created_at from tracks order by symbol, created_at desc;")

	return res, err
}
