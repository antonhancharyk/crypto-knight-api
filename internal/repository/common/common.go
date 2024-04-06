package common

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Common struct {
	db *sqlx.DB
}

func NewCommon(db *sqlx.DB) *Common {
	return &Common{db}
}

func (c *Common) On() {
	_, err := c.db.Exec("UPDATE common SET enabled = true WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Common) Off() {
	_, err := c.db.Exec("UPDATE common SET enabled = false WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}
}
