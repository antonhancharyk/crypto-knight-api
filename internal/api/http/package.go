package http

import (
	"fmt"
	"os"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
)

func RunHTTP(h *handler.Handler) {
	route.Init(h).Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
