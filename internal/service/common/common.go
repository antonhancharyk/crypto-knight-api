package common

import (
	"context"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/service/common"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
)

type Common struct {
	rep *repository.Repository
	pbCommon.UnimplementedCommonServiceServer
}

func NewCommon(rep *repository.Repository) *Common {
	return &Common{rep: rep}
}

func (c *Common) On() {
	c.rep.On()
}

func (c *Common) Off() {
	c.rep.Off()
}

func (c *Common) GetStatus(ctx context.Context, req *pbCommon.EmptyRequest) (*pbCommon.Enabled, error) {
	return &pbCommon.Enabled{Enabled: true}, nil
}
