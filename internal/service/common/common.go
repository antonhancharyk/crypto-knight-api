package common

import (
	"context"
	"log"

	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
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

func (c *Common) GetStatus() bool {
	return c.repo.GetStatus()
}

func (c *Common) On() {
	c.repo.On()
	_, err := grpc.Get().Common.SwitchOn(context.Background(), &pbCommon.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func (c *Common) Off() {
	c.repo.Off()
	_, err := grpc.Get().Common.SwitchOff(context.Background(), &pbCommon.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
