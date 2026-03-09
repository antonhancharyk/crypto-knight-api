package balance

import (
	"encoding/json"
	"testing"

	"github.com/antongoncharik/crypto-knight-api/internal/entity/balance"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
)

type stubClient struct {
	getResp []byte
	getErr  error
}

func (s *stubClient) Get(url string, isBot bool) ([]byte, error) { return s.getResp, s.getErr }
func (s *stubClient) Post(url string, body []byte, isBot bool) ([]byte, error) {
	return nil, nil
}
func (s *stubClient) Delete(url string, isBot bool) ([]byte, error) { return nil, nil }

func TestBalanceService_Get_Success(t *testing.T) {
	balances := []balance.Balance{
		{Asset: "BTC"},
		{Asset: "USDT", Balance: "100"},
	}
	data, _ := json.Marshal(balances)

	client := &stubClient{getResp: data}
	svc := New(client)

	res, err := svc.Get()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Asset != "USDT" || res.Balance != "100" {
		t.Fatalf("unexpected result: %+v", res)
	}
}

func TestBalanceService_Get_NotFound(t *testing.T) {
	balances := []balance.Balance{{Asset: "BTC"}}
	data, _ := json.Marshal(balances)

	client := &stubClient{getResp: data}
	svc := New(client)

	_, err := svc.Get()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestBalanceService_New(t *testing.T) {
	if New(&stubClient{}) == nil {
		t.Fatal("expected non-nil service")
	}
}

var _ api.Client = (*stubClient)(nil)

