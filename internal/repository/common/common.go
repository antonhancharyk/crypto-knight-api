package common

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type CommonData struct {
	Enabled bool `db:"enabled"`
}

type Common struct {
	db *sqlx.DB
}

func NewCommon(db *sqlx.DB) *Common {
	return &Common{db: db}
}

func (c *Common) GetStatus() bool {
	var commonData []CommonData

	err := c.db.Select(&commonData, "select enabled from common where id = 1")
	if err != nil {
		log.Fatal(err)
	}

	return commonData[0].Enabled
}

func (c *Common) On() {
	_, err := c.db.Exec("update common set enabled = true where id = 1")
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Common) Off() {
	_, err := c.db.Exec("update common set enabled = false where id = 1")
	if err != nil {
		log.Fatal(err)
	}
}
