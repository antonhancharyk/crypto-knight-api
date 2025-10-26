package tracks

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	tracks := req.Group("/tracks")

	tracks.GET("", hdl.Tracks.GetAll)
	tracks.POST("", hdl.Tracks.Create)
	tracks.POST("/bulk", hdl.Tracks.CreateBulk)
}

func UseRoutesHistory(req *gin.Engine, hdl *handler.Handler) {
	tracks := req.Group("/tracks/history")

	tracks.GET("", hdl.Tracks.GetAllHistory)
	tracks.POST("/bulk", hdl.Tracks.CreateBulkHistory)
}
