package tracks

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Tracks struct {
	svc *service.Service
}

func New(svc *service.Service) *Tracks {
	return &Tracks{svc}
}

func (t *Tracks) GetAll(ctx *gin.Context) {
	res := t.svc.Tracks.GetAll()

	ctx.JSON(http.StatusOK, res)
}

func (t *Tracks) Create(ctx *gin.Context) {
	var trackData track.Track

	err := ctx.ShouldBindJSON(&trackData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid JSON"})
		return
	}

	err = validate.Struct(trackData)
	if err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Tag()
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	t.svc.Tracks.Create(trackData)

	ctx.Status(http.StatusCreated)
}
