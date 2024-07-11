package app

import (
	"log"

	"github.com/antongoncharik/crypto-knight-api/internal/api/grpc"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/config"
	"github.com/antongoncharik/crypto-knight-api/internal/database"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
)

func Run() {
	keys, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	database.Connect()
	defer database.Close()

	grpc.Connect()
	defer grpc.Close()

	repo := repository.New(database.Get())
	svc := service.New(repo, keys)
	hdl := handler.New(svc)

	http.RunHTTP(hdl, keys)
}
