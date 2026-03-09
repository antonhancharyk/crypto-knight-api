package balance

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	entity "github.com/antongoncharik/crypto-knight-api/internal/entity/balance"
	"github.com/gin-gonic/gin"
)

type stubBalanceService struct {
	result entity.Balance
	err    error
}

func (s *stubBalanceService) Get() (entity.Balance, error) {
	return s.result, s.err
}

func TestBalanceHandler_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubBalanceService{
		result: entity.Balance{Asset: "USDT"},
	}

	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/balance", nil)
	ctx.Request = req

	h.Get(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var res entity.Balance
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if res.Asset != "USDT" {
		t.Fatalf("unexpected response body: %+v", res)
	}
}

