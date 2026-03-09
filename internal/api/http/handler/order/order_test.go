package order

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/order"
	"github.com/gin-gonic/gin"
)

type stubOrderService struct {
	result []order.Order
	err    error
}

func (s *stubOrderService) GetOpenOrders() ([]order.Order, error) {
	return s.result, s.err
}

func TestOrderHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubOrderService{
		result: []order.Order{{Symbol: "BTCUSDT"}},
	}

	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	ctx.Request = req

	h.GetAll(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var res []order.Order
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(res) != 1 || res[0].Symbol != "BTCUSDT" {
		t.Fatalf("unexpected response body: %+v", res)
	}
}

