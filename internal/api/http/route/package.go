package route

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route/common"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	common.CommonRoutes(router)

	return router
}
