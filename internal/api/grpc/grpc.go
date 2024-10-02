package grpc

import (
	"log"

	pbCommon "github.com/antongoncharik/crypto-knight-protos/gen/go/common"
	"google.golang.org/grpc"
)

type GRPCClients struct {
	Common pbCommon.CommonServiceClient
}

var gRPCClients *GRPCClients
var clientConn *grpc.ClientConn

func Connect() (*grpc.ClientConn, *GRPCClients) {
	conn, err := grpc.Dial("113.30.189.245:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	clientConn = conn
	gRPCClients = &GRPCClients{Common: pbCommon.NewCommonServiceClient(clientConn)}

	log.Println("gRPS client is running")

	return clientConn, gRPCClients
}
