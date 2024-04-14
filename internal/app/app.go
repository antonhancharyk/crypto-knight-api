package app

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/database"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
)

func Run() {
	database.Connect()
	defer database.Close()

	db := database.Get()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	hdl := handler.NewHandler(svc)

	go grpc.RunGRPC(svc)
	go http.RunHTTP(hdl)

	select {}
}
