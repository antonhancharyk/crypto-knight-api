package common

import (
	"context"
	"log"

	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	pbCommon "github.com/antongoncharik/crypto-knight-protos/gen/go/common"
)

type Common struct {
	repo *repository.Repository
	pbCommon.UnimplementedCommonServiceServer
}

func New(repo *repository.Repository) *Common {
	return &Common{repo: repo}
}

func (c *Common) GetStatus() bool {
	return c.repo.Common.GetStatus()
}

func (c *Common) Enable() {
	c.repo.Common.On()
	_, err := grpc.Get().Common.Enable(context.Background(), &pbCommon.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func (c *Common) Disable() {
	c.repo.Common.Off()
	_, err := grpc.Get().Common.Disable(context.Background(), &pbCommon.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
