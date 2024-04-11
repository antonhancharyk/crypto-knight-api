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

	d := database.Get()
	r := repository.NewRepository(d)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	grpc.RunGRPC(s)
	http.RunHTTP(h)
}
