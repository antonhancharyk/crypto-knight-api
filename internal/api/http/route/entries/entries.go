package entries

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(req *gin.Engine, hdl *handler.Handler) {
	entries := req.Group("/entries")

	entries.GET("", hdl.Entries.GetAll)
	entries.POST("", hdl.Entries.Create)
}
