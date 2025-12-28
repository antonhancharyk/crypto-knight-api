package http

import (
	"log"
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
)

func RunHTTP(hdl *handler.Handler, keys auth.RSAKeys) *http.Server {
	router := route.Init(hdl, keys)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Println("HTTP server is running")

	return srv
}
