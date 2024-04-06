package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
)

type Common interface {
	On()
	Off()
}

type Service struct {
	Common
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Common: common.NewCommon(rep),
	}
}
