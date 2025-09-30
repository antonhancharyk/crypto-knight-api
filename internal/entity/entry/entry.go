package entry

import (
	"time"

	"github.com/lib/pq"
)

type Entry struct {
	Symbol     string          `json:"symbol" db:"symbol" validate:"required"`
	HighPrice  float64         `json:"high_price" db:"high_price" validate:"numeric"`
	LowPrice   float64         `json:"low_price" db:"low_price" validate:"numeric"`
	HighPrices pq.Float64Array `json:"high_prices" db:"high_prices"`
	LowPrices  pq.Float64Array `json:"low_prices" db:"low_prices"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" db:"updated_at"`
}

type QueryParams struct {
	From   string `form:"from"`
	To     string `form:"to"`
	Symbol string `form:"symbol"`
}
