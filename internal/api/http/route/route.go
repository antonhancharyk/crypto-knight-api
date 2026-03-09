package route

import (
	"net/http"

	_ "github.com/antongoncharik/crypto-knight-api/docs"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/cors"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/logging"
	authRoutes "github.com/antongoncharik/crypto-knight-api/internal/api/http/route/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/balance"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/common"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/entries"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/klines"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/order"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/position"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/tracks"
	authEntity "github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(hdl *handler.Handler, keys authEntity.RSAKeys) *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", healthz)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(logging.UseLogging())
	router.Use(cors.UseCORS())
	router.Use(auth.UseAuth(keys))

	common.UseRoutes(router, hdl)
	tracks.UseRoutes(router, hdl)
	tracks.UseRoutesHistory(router, hdl)
	authRoutes.UseRoutes(router, hdl)
	entries.UseRoutes(router, hdl)
	position.UseRoutes(router, hdl)
	balance.UseRoutes(router, hdl)
	order.UseRoutes(router, hdl)
	klines.UseRoutes(router, hdl)

	return router
}

func healthz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
