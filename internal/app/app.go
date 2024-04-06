package app

import (
	"fmt"
	"os"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
	"github.com/antongoncharik/crypto-knight-api/internal/database"
	"github.com/antongoncharik/crypto-knight-api/internal/repository"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
)

func Run() {
	database.Connect()
	defer database.Close()

	r := repository.NewRepository(database.Get())
	s := service.NewService(r)
	h := handler.NewHandler(s)

	route.Init(h).Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
