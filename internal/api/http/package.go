package http

import (
	"fmt"
	"os"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
)

func RunHTTP(hdl *handler.Handler) {
	route.Init(hdl).Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
