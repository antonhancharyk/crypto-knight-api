package order

import (
	"encoding/json"
	"testing"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/order"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
)

type stubOrderClient struct {
	getResp []byte
	getErr  error
}

func (s *stubOrderClient) Get(url string, isBot bool) ([]byte, error) { return s.getResp, s.getErr }
func (s *stubOrderClient) Post(url string, body []byte, isBot bool) ([]byte, error) {
	return nil, nil
}
func (s *stubOrderClient) Delete(url string, isBot bool) ([]byte, error) { return nil, nil }

func TestOrderService_GetOpenOrders_Success(t *testing.T) {
	orders := []order.Order{{Symbol: "BTCUSDT"}}
	data, _ := json.Marshal(orders)

	client := &stubOrderClient{getResp: data}
	svc := New(client)

	res, err := svc.GetOpenOrders()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res) != 1 || res[0].Symbol != "BTCUSDT" {
		t.Fatalf("unexpected result: %+v", res)
	}
}

func TestOrderService_New(t *testing.T) {
	if New(&stubOrderClient{}) == nil {
		t.Fatal("expected non-nil service")
	}
}

var _ api.Client = (*stubOrderClient)(nil)

