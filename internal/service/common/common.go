package common

import (
	"context"

	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	pbCommon "github.com/antongoncharik/crypto-knight-protos/gen/go/common"
)

type Common struct {
	repo        *repository.Repository
	grpcClients *grpc.GRPCClients
	pbCommon.UnimplementedCommonServiceServer
}

func New(repo *repository.Repository, grpcClients *grpc.GRPCClients) *Common {
	return &Common{repo: repo, grpcClients: grpcClients}
}

func (c *Common) GetStatus() (bool, error) {
	return c.repo.Common.GetStatus()
}

func (c *Common) Enable() error {
	err := c.repo.Common.On()
	if err != nil {
		return err
	}

	_, err = c.grpcClients.Common.Enable(context.Background(), &pbCommon.EmptyRequest{})

	return err
}

func (c *Common) Disable() error {
	err := c.repo.Common.Off()
	if err != nil {
		return err
	}

	_, err = c.grpcClients.Common.Disable(context.Background(), &pbCommon.EmptyRequest{})

	return err
}
