package route

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/middleware/cors"
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/common"
	"github.com/gin-gonic/gin"
)

func Init(hdl *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.UseCORS())

	common.UseCommonRoutes(router, hdl)

	return router
}
