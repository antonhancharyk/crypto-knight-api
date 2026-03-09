package response

import (
	"net/http"

	"github.com/antongoncharik/crypto-knight-api/internal/errors"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error any `json:"error"`
}

func JSON(ctx *gin.Context, status int, payload any) {
	ctx.JSON(status, payload)
}

func OK(ctx *gin.Context, payload any) {
	ctx.JSON(http.StatusOK, payload)
}

func Created(ctx *gin.Context, payload any) {
	ctx.JSON(http.StatusCreated, payload)
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func Error(ctx *gin.Context, status int, err error) {
	if err == nil {
		ctx.Status(status)
		return
	}

	ctx.JSON(status, ErrorResponse{Error: err.Error()})
}

func ValidationError(ctx *gin.Context, details any) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: details})
}

// WriteError writes err to ctx using domain error status when applicable, otherwise 500.
func WriteError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	status := errors.StatusCode(err)
	msg := errors.Message(err)
	ctx.JSON(status, ErrorResponse{Error: msg})
}
