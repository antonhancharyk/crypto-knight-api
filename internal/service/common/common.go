package common

import "github.com/antongoncharik/crypto-knight-api/internal/repository"

type Common struct {
	rep *repository.Repository
}

func NewCommon(rep *repository.Repository) *Common {
	return &Common{rep}
}

func (c *Common) On() {
	c.rep.On()
}

func (c *Common) Off() {
	c.rep.Off()
}
