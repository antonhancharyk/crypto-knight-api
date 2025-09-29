package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/config"
	"github.com/antongoncharik/crypto-knight-api/internal/database"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
	"github.com/joho/godotenv"
)

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	keys, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := database.Connect()
	defer db.Close()

	grpcClientConn, grpcClients := grpc.Connect()
	defer grpcClientConn.Close()

	apiClient := api.New()

	repo := repository.New(db)
	svc := service.New(repo, keys, grpcClients, apiClient)
	hdl := handler.New(svc)

	srv := http.RunHTTP(hdl, keys)

	<-quit

	log.Println("Shutting down HTTP server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down db...")

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down gRPC server...")

	err = grpcClientConn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
