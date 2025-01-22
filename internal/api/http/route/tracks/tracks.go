package tracks

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	commonGroup := req.Group("/tracks")

	commonGroup.GET("", hdl.Tracks.GetAll)
	commonGroup.POST("", hdl.Tracks.Create)
	commonGroup.POST("/bulk", hdl.Tracks.CreateBulk)
}
