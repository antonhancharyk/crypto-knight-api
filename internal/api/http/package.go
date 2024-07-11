package http

import (
	"fmt"
	"os"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
)

func RunHTTP(hdl *handler.Handler, keys auth.RSAKeys) {
	route.Init(hdl, keys).Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
