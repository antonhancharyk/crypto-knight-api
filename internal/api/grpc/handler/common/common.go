package common

import (
	"context"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/service/common"
)

type Server struct {
	pbCommon.UnimplementedCommonServiceServer
}

func (s *Server) GetStatus(ctx context.Context, req *pbCommon.EmptyRequest) (*pbCommon.Enabled, error) {
	return &pbCommon.Enabled{Enabled: true}, nil
}
