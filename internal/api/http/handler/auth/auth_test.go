package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type stubAuthService struct {
	err error
}

func (s *stubAuthService) ValidateToken(token string) error {
	return s.err
}

func TestAuthHandler_ValidateToken_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubAuthService{}
	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/auth/validate", nil)
	ctx.Request = req

	h.ValidateToken(ctx)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestAuthHandler_ValidateToken_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubAuthService{err: nil}
	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/auth/validate?token=abc", nil)
	ctx.Request = req

	h.ValidateToken(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

