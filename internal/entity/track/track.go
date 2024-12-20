package track

import (
	"time"

	"github.com/lib/pq"
)

type Track struct {
	Symbol           string          `json:"symbol" db:"symbol" validate:"required"`
	HighPrice        float64         `json:"high_price" db:"high_price" validate:"numeric"`
	LowPrice         float64         `json:"low_price" db:"low_price" validate:"numeric"`
	HighPrice1       float64         `json:"high_price_1" db:"high_price_1" validate:"numeric"`
	LowPrice1        float64         `json:"low_price_1" db:"low_price_1" validate:"numeric"`
	HighPrice2       float64         `json:"high_price_2" db:"high_price_2" validate:"numeric"`
	LowPrice2        float64         `json:"low_price_2" db:"low_price_2" validate:"numeric"`
	HighPrice3       float64         `json:"high_price_3" db:"high_price_3" validate:"numeric"`
	LowPrice3        float64         `json:"low_price_3" db:"low_price_3" validate:"numeric"`
	Causes           pq.StringArray  `json:"causes" db:"causes"`
	IsOrder          bool            `json:"is_order" db:"is_order"`
	HighCreatedAt    time.Time       `json:"high_created_at" db:"high_created_at"`
	LowCreatedAt     time.Time       `json:"low_created_at" db:"low_created_at"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at" db:"updated_at"`
	ResistancePrice1 float64         `json:"resistance_price_1" db:"resistance_price_1" `
	SupportPrice1    float64         `json:"support_price_1" db:"support_price_1" `
	ResistancePrice2 float64         `json:"resistance_price_2" db:"resistance_price_2" `
	SupportPrice2    float64         `json:"support_price_2" db:"support_price_2" `
	ResistancePrice3 float64         `json:"resistance_price_3" db:"resistance_price_3" `
	SupportPrice3    float64         `json:"support_price_3" db:"support_price_3" `
	HighPrices       pq.Float64Array `json:"high_prices" db:"high_prices"`
	LowPrices        pq.Float64Array `json:"low_prices" db:"low_prices"`
}

type QueryParams struct {
	From   string `form:"from"`
	To     string `form:"to"`
	Symbol string `form:"symbol"`
	Full   bool   `form:"full"`
}
