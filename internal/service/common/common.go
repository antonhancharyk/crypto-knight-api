package common

import (
	"context"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/pb/common"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
)

type Common struct {
	repo *repository.Repository
	pbCommon.UnimplementedCommonServiceServer
}

func NewCommon(repo *repository.Repository) *Common {
	return &Common{repo: repo}
}

func (c *Common) GetStatus(ctx context.Context, req *pbCommon.EmptyRequest) (*pbCommon.Enabled, error) {
	status := c.repo.GetStatus()
	return &pbCommon.Enabled{Enabled: status}, nil
}

func (c *Common) On() {
	c.repo.On()
}

func (c *Common) Off() {
	c.repo.Off()
}
