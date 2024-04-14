package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	pbCommon "github.com/antongoncharik/crypto-knight-api/internal/api/grpc/pb/common"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/antongoncharik/crypto-knight-api/internal/service/common"
	"google.golang.org/grpc"
)

func RunGRPC(s *service.Service) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_PORT_GRPC")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	pbCommon.RegisterCommonServiceServer(srv, s.Common.(*common.Common))

	log.Printf("Server listening on port %s", os.Getenv("APP_PORT_GRPC"))

	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
