package service

import (
	"context"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/pb/common"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
)

type Common interface {
	On()
	Off()
	GetStatus(ctx context.Context, req *pbCommon.EmptyRequest) (*pbCommon.Enabled, error)
}

type Service struct {
	Common
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Common: common.NewCommon(repo),
	}
}
