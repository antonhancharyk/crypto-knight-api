package tracks

import (
	"net/http"
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/api/http/response"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/track"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TracksService interface {
	GetAll(queryParams track.QueryParams) ([]track.Track, error)
	Create(track track.Track) error
	CreateBulk(tracks []track.Track) error
	GetAllHistory(queryParams track.QueryParams) ([]track.Track, error)
	CreateBulkHistory(tracks []track.Track) error
	GetLastTracks() ([]track.Track, error)
}

type Tracks struct {
	svc TracksService
}

func New(svc TracksService) *Tracks {
	return &Tracks{svc: svc}
}

// GetAll godoc
// @Summary      List tracks
// @Tags         tracks
// @Param        from   query  string  false  "From (YYYY-MM-DD HH:MI:SS)"
// @Param        to     query  string  false  "To (YYYY-MM-DD HH:MI:SS)"
// @Param        symbol query  string  false  "Symbol"
// @Produce      json
// @Success      200  {array}   track.Track
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /tracks [get]
func (t *Tracks) GetAll(ctx *gin.Context) {
	var queryParams track.QueryParams

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	now := time.Now()

	if queryParams.From == "" {
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		queryParams.From = startOfDay.Format("2006-01-02 15:04:05")
	}
	if queryParams.To == "" {
		endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		queryParams.To = endOfDay.Format("2006-01-02 15:04:05")
	}

	if _, err := time.Parse("2006-01-02 15:04:05", queryParams.From); err != nil {
		response.ValidationError(ctx, "Invalid date format. Expected 'YYYY-MM-DD HH:MI:SS'")
		return
	}
	if _, err := time.Parse("2006-01-02 15:04:05", queryParams.To); err != nil {
		response.ValidationError(ctx, "Invalid date format. Expected 'YYYY-MM-DD HH:MI:SS'")
		return
	}

	res, err := t.svc.GetAll(queryParams)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.OK(ctx, res)
}

// Create godoc
// @Summary      Create track
// @Tags         tracks
// @Accept       json
// @Param        body  body  track.Track  true  "Track"
// @Success      201
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /tracks [post]
func (t *Tracks) Create(ctx *gin.Context) {
	var trackData track.Track

	if err := ctx.ShouldBindJSON(&trackData); err != nil {
		response.ValidationError(ctx, "invalid JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(trackData); err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Tag()
		}
		response.ValidationError(ctx, errors)
		return
	}

	if err := t.svc.Create(trackData); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

// CreateBulk godoc
// @Summary      Create tracks in bulk
// @Tags         tracks
// @Accept       json
// @Param        body  body  array  true  "Tracks"
// @Success      201
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /tracks/bulk [post]
func (t *Tracks) CreateBulk(ctx *gin.Context) {
	var tracksData []track.Track

	if err := ctx.ShouldBindJSON(&tracksData); err != nil {
		response.ValidationError(ctx, "invalid JSON")
		return
	}

	validate := validator.New()

	errors := make(map[int]map[string]string)
	for i, track := range tracksData {
		if err := validate.Struct(track); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				fieldErrors := make(map[string]string)
				for _, e := range validationErrors {
					fieldErrors[e.Field()] = e.Tag()
				}
				errors[i] = fieldErrors
			} else {
				response.ValidationError(ctx, "Validation error occurred")
				return
			}
		}
	}

	if len(errors) != 0 {
		response.ValidationError(ctx, errors)
		return
	}

	if err := t.svc.CreateBulk(tracksData); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

// GetLastTracks godoc
// @Summary      Get last tracks
// @Tags         tracks
// @Produce      json
// @Success      200  {array}   track.Track
// @Failure      500  {object}  response.ErrorResponse
// @Router       /tracks/last [get]
func (t *Tracks) GetLastTracks(ctx *gin.Context) {
	res, err := t.svc.GetLastTracks()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.OK(ctx, res)
}

// GetAllHistory godoc
// @Summary      List tracks history
// @Tags         tracks-history
// @Param        from   query  string  false  "From (YYYY-MM-DD HH:MI:SS)"
// @Param        to     query  string  false  "To (YYYY-MM-DD HH:MI:SS)"
// @Param        symbol query  string  false  "Symbol"
// @Produce      json
// @Success      200  {array}   track.Track
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /tracks/history [get]
func (t *Tracks) GetAllHistory(ctx *gin.Context) {
	var queryParams track.QueryParams

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	now := time.Now()

	if queryParams.From == "" {
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		queryParams.From = startOfDay.Format("2006-01-02 15:04:05")
	}
	if queryParams.To == "" {
		endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		queryParams.To = endOfDay.Format("2006-01-02 15:04:05")
	}

	if _, err := time.Parse("2006-01-02 15:04:05", queryParams.From); err != nil {
		response.ValidationError(ctx, "Invalid date format. Expected 'YYYY-MM-DD HH:MI:SS'")
		return
	}
	if _, err := time.Parse("2006-01-02 15:04:05", queryParams.To); err != nil {
		response.ValidationError(ctx, "Invalid date format. Expected 'YYYY-MM-DD HH:MI:SS'")
		return
	}

	res, err := t.svc.GetAllHistory(queryParams)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.OK(ctx, res)
}

// CreateBulkHistory godoc
// @Summary      Create tracks history in bulk
// @Tags         tracks-history
// @Accept       json
// @Param        body  body  array  true  "Tracks"
// @Success      201
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /tracks/history/bulk [post]
func (t *Tracks) CreateBulkHistory(ctx *gin.Context) {
	var tracksData []track.Track

	if err := ctx.ShouldBindJSON(&tracksData); err != nil {
		response.ValidationError(ctx, "invalid JSON")
		return
	}

	validate := validator.New()

	errors := make(map[int]map[string]string)
	for i, track := range tracksData {
		if err := validate.Struct(track); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				fieldErrors := make(map[string]string)
				for _, e := range validationErrors {
					fieldErrors[e.Field()] = e.Tag()
				}
				errors[i] = fieldErrors
			} else {
				response.ValidationError(ctx, "Validation error occurred")
				return
			}
		}
	}

	if len(errors) != 0 {
		response.ValidationError(ctx, errors)
		return
	}

	if err := t.svc.CreateBulkHistory(tracksData); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
}
