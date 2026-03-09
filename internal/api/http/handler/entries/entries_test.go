package entries

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/entry"
	"github.com/gin-gonic/gin"
)

type stubEntriesService struct {
	result []entry.Entry
	err    error
}

func (s *stubEntriesService) GetAll() ([]entry.Entry, error) {
	return s.result, s.err
}

func (s *stubEntriesService) Create(e entry.Entry) error {
	return s.err
}

func TestEntriesHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubEntriesService{
		result: []entry.Entry{{Symbol: "BTCUSDT"}},
	}

	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodGet, "/entries", nil)
	ctx.Request = req

	h.GetAll(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var res []entry.Entry
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(res) != 1 || res[0].Symbol != "BTCUSDT" {
		t.Fatalf("unexpected response body: %+v", res)
	}
}

func TestEntriesHandler_Create_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubEntriesService{}
	h := New(stubSvc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodPost, "/entries", bytes.NewBufferString("not-json"))
	ctx.Request = req

	h.Create(ctx)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestEntriesHandler_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	stubSvc := &stubEntriesService{}
	h := New(stubSvc)

	body := `{"symbol":"BTCUSDT","high_price":1.0,"low_price":0.5}`
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodPost, "/entries", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req

	h.Create(ctx)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

