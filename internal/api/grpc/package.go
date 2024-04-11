package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/service/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
	"google.golang.org/grpc"
)

func RunGRPC(s *service.Service) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_PORT_GRPC")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ser := grpc.NewServer()

	pbCommon.RegisterCommonServiceServer(ser, s.Common.(*common.Common))

	log.Println("Server listening on port 50051")

	err = ser.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
