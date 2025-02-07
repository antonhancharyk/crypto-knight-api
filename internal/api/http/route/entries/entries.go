package entries

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	commonGroup := req.Group("/entries")

	commonGroup.GET("", hdl.Entries.GetAll)
	commonGroup.POST("", hdl.Entries.Create)
}
