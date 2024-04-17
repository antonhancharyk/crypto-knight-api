package grpc

import (
	"log"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/pb/common"
	"google.golang.org/grpc"
)

type GRPCClients struct {
	Common pbCommon.CommonServiceClient
}

var gRPCClients *GRPCClients
var clientConn *grpc.ClientConn

func Connect() {
	conn, err := grpc.Dial("113.30.189.245:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	clientConn = conn
	gRPCClients = &GRPCClients{Common: pbCommon.NewCommonServiceClient(clientConn)}
}

func Get() *GRPCClients {
	return gRPCClients
}

func Close() {
	clientConn.Close()
}
