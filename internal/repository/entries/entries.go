package entries

import (
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Entries struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Entries {
	return &Entries{db: db}
}

func (t *Entries) GetAll() ([]entry.Entry, error) {
	res := []entry.Entry{}
	err := t.db.Select(&res, "select distinct on (symbol) symbol, high_price, low_price, created_at from last_entries order by symbol, created_at desc;")

	return res, err
}

func (t *Entries) Create(entry entry.Entry) error {
	_, err := t.db.Exec(`insert into last_entries (symbol, high_price, low_price) values ($1, $2, $3)`, entry.Symbol, entry.HighPrice, entry.LowPrice)

	return err
}
