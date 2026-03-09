package http

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
	"github.com/antongoncharik/crypto-knight-api/internal/config"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/logger"
)

func RunHTTP(hdl *handler.Handler, keys auth.RSAKeys, httpCfg config.HTTPConfig) *http.Server {
	router := route.Init(hdl, keys)
	srv := &http.Server{
		Addr:    httpCfg.Addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatalw("HTTP server failed", "error", err)
		}
	}()

	logger.Log.Infow("HTTP server is running", "addr", httpCfg.Addr)

	return srv
}
