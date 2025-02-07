package route

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/cors"
	authRoutes "github.com/antongoncharik/crypto-knight-api/internal/api/http/route/auth"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/common"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/entries"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/tracks"
	authEntity "github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/gin-gonic/gin"
)

func Init(hdl *handler.Handler, keys authEntity.RSAKeys) *gin.Engine {
	router := gin.Default()

	router.Use(cors.UseCORS())
	router.Use(auth.UseAuth(keys))

	common.UseRoutes(router, hdl)
	tracks.UseRoutes(router, hdl)
	authRoutes.UseRoutes(router, hdl)
	entries.UseRoutes(router, hdl)

	return router
}
