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
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, high_prices, low_prices, created_at from tracks where created_at between $1 and $2 and symbol = $3 order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol)

		return tracksData, err
	}

	if queryParams.Full && queryParams.Symbol == "" {
		err := t.db.Select(&tracksData, "select symbol, high_price, low_price, high_prices, low_prices, created_at from tracks where created_at between $1 and $2 order by created_at desc", queryParams.From, queryParams.To)

		return tracksData, err
	}

	if queryParams.Symbol != "" {
		err := t.db.Select(&tracksData, "with ranked_tracks as (select id, symbol, high_price, low_price, created_at, lag(high_price) over (partition by symbol order by created_at) as prev_high_price, lag(low_price) over (partition by symbol order by created_at) as prev_low_price, high_prices, low_prices from tracks where created_at between $1 and $2 and symbol = $3) select symbol, high_price, low_price, created_at, high_prices, low_prices from ranked_tracks where high_price != prev_high_price or low_price != prev_low_price or prev_high_price is null order by created_at desc", queryParams.From, queryParams.To, queryParams.Symbol)

		return tracksData, err
	}

	err := t.db.Select(&tracksData, "with ranked_tracks as (select id, symbol, high_price, low_price, created_at, lag(high_price) over (partition by symbol order by created_at) as prev_high_price, lag(low_price) over (partition by symbol order by created_at) as prev_low_price, high_prices, low_prices from tracks where created_at between $1 and $2) select symbol, high_price, low_price, created_at, high_prices, low_prices from ranked_tracks where high_price != prev_high_price or low_price != prev_low_price or prev_high_price is null order by created_at desc", queryParams.From, queryParams.To)

	return tracksData, err
}

func (t *Tracks) Create(track track.Track) error {
	var err error
	if (track.CreatedAt == time.Time{}) {
		_, err = t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price, high_prices, low_prices)
		VALUES ($1, $2, $3, $4, $5)`, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrices, track.LowPrices)
	} else {
		_, err = t.db.Exec(`INSERT INTO tracks (symbol, high_price, low_price, high_prices, low_prices, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrices, track.LowPrices, track.CreatedAt)
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

func (t *Tracks) GetAllHistory(q track.QueryParams) ([]track.Track, error) {
	var (
		tracks []track.Track
		args   []any
	)

	baseSelect := `
		select
			symbol,
			high_price,
			low_price,
			created_at,
			high_prices,
			low_prices,
			middle_price_high,
			middle_price_low
	`

	builder := strings.Builder{}

	if q.Full {
		builder.WriteString(baseSelect)
		builder.WriteString(" from tracks_history where created_at between $1 and $2")
		args = append(args, q.From, q.To)

		argPos := 3

		if q.Symbol != "" {
			fmt.Fprintf(&builder, " and symbol = $%d", argPos)
			args = append(args, q.Symbol)
			argPos++
		}

		fmt.Fprintf(&builder, " and interval = $%d", argPos)
		args = append(args, q.Interval)
		argPos++

		if q.ShowOnlyEntries {
			builder.WriteString(" and (high_price = 1 or low_price = 1)")
		}

		builder.WriteString(" order by created_at desc")

		err := t.db.Select(&tracks, builder.String(), args...)
		return tracks, err
	}

	builder.WriteString(`
		with ranked_tracks as (
			select
				id,
				symbol,
				high_price,
				low_price,
				created_at,
				lag(high_price) over (partition by symbol order by created_at) as prev_high_price,
				lag(low_price)  over (partition by symbol order by created_at) as prev_low_price,
				high_prices,
				low_prices,
				middle_price_high,
				middle_price_low
			from tracks_history
			where created_at between $1 and $2
	`)

	args = append(args, q.From, q.To)
	argPos := 3

	if q.Symbol != "" {
		fmt.Fprintf(&builder, " and symbol = $%d", argPos)
		args = append(args, q.Symbol)
		argPos++
	}

	fmt.Fprintf(&builder, " and interval = $%d", argPos)
	args = append(args, q.Interval)
	argPos++

	if q.ShowOnlyEntries {
		builder.WriteString(" and (high_price = 1 or low_price = 1)")
	}

	builder.WriteString(`
	)
	`)

	builder.WriteString(baseSelect)
	builder.WriteString(`
		from ranked_tracks
		where
			prev_high_price is null
			or high_price != prev_high_price
			or low_price  != prev_low_price
		order by created_at desc
	`)

	err := t.db.Select(&tracks, builder.String(), args...)
	return tracks, err
}

func (t *Tracks) CreateBulkHistory(tracks []track.Track) error {
	var placeholders []string
	var values []any
	for i, track := range tracks {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*9+1, i*9+2, i*9+3, i*9+4, i*9+5, i*9+6, i*9+7, i*9+8, i*9+9))
		values = append(values, track.Symbol, track.HighPrice, track.LowPrice, track.HighPrices, track.LowPrices, track.CreatedAt, track.Interval, track.MiddlePriceHigh, track.MiddlePriceLow)
	}

	query := fmt.Sprintf(`
		INSERT INTO tracks_history (symbol, high_price, low_price, high_prices, low_prices, created_at, interval, middle_price_high, middle_price_low)
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
