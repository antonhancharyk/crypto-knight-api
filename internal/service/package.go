package service

import (
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
)

type Common interface {
	Enable()
	Disable()
	GetStatus() bool
}

type Service struct {
	Common
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Common: common.NewCommon(repo),
	}
}
