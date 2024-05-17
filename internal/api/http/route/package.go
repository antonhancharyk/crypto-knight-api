package route

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/cors"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/common"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/tracks"
	"github.com/gin-gonic/gin"
)

func Init(hdl *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.UseCORS())

	common.UseRoutes(router, hdl)
	tracks.UseRoutes(router, hdl)

	return router
}
