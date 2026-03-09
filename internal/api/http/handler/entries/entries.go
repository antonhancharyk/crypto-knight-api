package entries

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EntriesService interface {
	GetAll() ([]entry.Entry, error)
	Create(entry entry.Entry) error
}

type Entries struct {
	svc EntriesService
}

func New(svc EntriesService) *Entries {
	return &Entries{svc: svc}
}

// GetAll godoc
// @Summary      List all entries
// @Tags         entries
// @Produce      json
// @Success      200  {array}   entry.Entry
// @Failure      500  {object}  response.ErrorResponse
// @Router       /entries [get]
func (t *Entries) GetAll(ctx *gin.Context) {
	res, err := t.svc.GetAll()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.OK(ctx, res)
}

// Create godoc
// @Summary      Create entry
// @Tags         entries
// @Accept       json
// @Param        body  body  entry.Entry  true  "Entry"
// @Success      201
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /entries [post]
func (t *Entries) Create(ctx *gin.Context) {
	var entry entry.Entry

	if err := ctx.ShouldBindJSON(&entry); err != nil {
		response.ValidationError(ctx, "invalid JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(entry); err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Tag()
		}
		response.ValidationError(ctx, errors)
		return
	}

	if err := t.svc.Create(entry); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.Created(ctx, nil)
}
