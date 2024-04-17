package grpc

import (
	"context"
	"fmt"
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
	fmt.Println("HELLO 0")
	conn, err := grpc.Dial("113.30.189.245:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("HELLO 1")

	clientConn = conn
	gRPCClients = &GRPCClients{Common: pbCommon.NewCommonServiceClient(clientConn)}
	fmt.Println("HELLO 2")
	Get().Common.SwitchOn(context.Background(), &pbCommon.EmptyRequest{})
	fmt.Println("HELLO 3")
}

func Get() *GRPCClients {
	return gRPCClients
}

func Close() {
	clientConn.Close()
}
