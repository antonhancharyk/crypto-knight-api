package grpc

import (
	"github.com/antongoncharik/crypto-knight-api/internal/logger"
	pbCommon "github.com/antongoncharik/crypto-knight-protos/gen/go/common"
	"google.golang.org/grpc"
)

type GRPCClients struct {
	Common pbCommon.CommonServiceClient
}

func Connect(host string) (*grpc.ClientConn, *GRPCClients, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	clients := &GRPCClients{
		Common: pbCommon.NewCommonServiceClient(conn),
	}

	logger.Log.Infow("gRPC client is running", "host", host)

	return conn, clients, nil
}
