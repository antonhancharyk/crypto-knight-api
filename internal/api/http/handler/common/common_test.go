package common

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type stubCommonService struct {
	status bool
	err    error
}

func (s *stubCommonService) GetStatus() (bool, error) {
	return s.status, s.err
}

func (s *stubCommonService) Enable() error {
	return s.err
}

func (s *stubCommonService) Disable() error {
	return s.err
}

func TestCommonHandler_GetStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubCommonService{status: true}
	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/common/status", nil)
	ctx.Request = req

	h.GetStatus(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestCommonHandler_GetStatus_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubCommonService{err: errors.New("test error")}
	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/common/status", nil)
	ctx.Request = req

	h.GetStatus(ctx)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

