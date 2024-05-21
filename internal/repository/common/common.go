package common

import (
	"github.com/jmoiron/sqlx"
)

type CommonData struct {
	Enabled bool `db:"enabled"`
}

type Common struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Common {
	return &Common{db: db}
}

func (c *Common) GetStatus() (bool, error) {
	var commonData []CommonData

	err := c.db.Select(&commonData, "select enabled from common where id = 1")

	return commonData[0].Enabled, err
}

func (c *Common) On() error {
	_, err := c.db.Exec("update common set enabled = true where id = 1")

	return err
}

func (c *Common) Off() error {
	_, err := c.db.Exec("update common set enabled = false where id = 1")

	return err
}
