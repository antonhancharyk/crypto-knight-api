package entries

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/antongoncharik/crypto-knight-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Entries struct {
	svc *service.Service
}

func New(svc *service.Service) *Entries {
	return &Entries{svc}
}

func (t *Entries) GetAll(ctx *gin.Context) {
	res, err := t.svc.Entries.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (t *Entries) Create(ctx *gin.Context) {
	var entry entry.Entry

	err := ctx.ShouldBindJSON(&entry)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	validate := validator.New()
	err = validate.Struct(entry)
	if err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Tag()
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	err = t.svc.Entries.Create(entry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
