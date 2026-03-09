package main

import (
	"github.com/antongoncharik/crypto-knight-api/internal/app"
)

// @title           Crypto Knight API
// @version         1.0
// @description     API for crypto-knight trading (tracks, entries, balance, orders, klines).
// @host            localhost:8080
// @BasePath        /
// @schemes         http

func main() {
	app.Run()
}
