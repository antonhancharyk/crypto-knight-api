package track

import (
	"time"

	"github.com/lib/pq"
)

type Track struct {
	Symbol        string          `json:"symbol" db:"symbol" validate:"required"`
	HighPrice     float64         `json:"high_price" db:"high_price" validate:"numeric"`
	LowPrice      float64         `json:"low_price" db:"low_price" validate:"numeric"`
	IsOrder       bool            `json:"is_order" db:"is_order"`
	HighCreatedAt time.Time       `json:"high_created_at" db:"high_created_at"`
	LowCreatedAt  time.Time       `json:"low_created_at" db:"low_created_at"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" db:"updated_at"`
	HighPrices    pq.Float64Array `json:"high_prices" db:"high_prices"`
	LowPrices     pq.Float64Array `json:"low_prices" db:"low_prices"`
}

type QueryParams struct {
	From   string `form:"from"`
	To     string `form:"to"`
	Symbol string `form:"symbol"`
	Full   bool   `form:"full"`
}
