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
	"github.com/antongoncharik/crypto-knight-api/internal/logger"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
	"github.com/joho/godotenv"
)

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	if err := logger.Init(cfg.Env); err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	db, err := database.Connect(cfg.DB)
	if err != nil {
		logger.Log.Fatalw("failed to connect to database", "error", err)
	}

	grpcClientConn, grpcClients, err := grpc.Connect(cfg.GRPC.Host)
	if err != nil {
		logger.Log.Fatalw("failed to connect to gRPC", "error", err)
	}

	apiClient := api.New()

	repo := repository.New(db)
	svc := service.New(repo, cfg.Auth, grpcClients, apiClient)
	hdl := handler.New(svc)

	srv := http.RunHTTP(hdl, cfg.Auth, cfg.HTTP)

	<-quit

	logger.Log.Info("Shutting down HTTP server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatalw("failed to shutdown HTTP server", "error", err)
	}

	logger.Log.Info("Shutting down db...")

	if err := db.Close(); err != nil {
		logger.Log.Fatalw("failed to close db", "error", err)
	}

	logger.Log.Info("Shutting down gRPC server...")

	if err := grpcClientConn.Close(); err != nil {
		logger.Log.Fatalw("failed to close gRPC connection", "error", err)
	}
}
