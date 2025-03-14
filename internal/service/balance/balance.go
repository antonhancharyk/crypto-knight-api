package balance

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/constant"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/balance"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
	"github.com/elliotchance/pie/v2"
)

type BalanceService struct {
	api *api.HTTPClient
}

func New(api *api.HTTPClient) *BalanceService {
	return &BalanceService{api: api}
}

func (b *BalanceService) Get() (balance.Balance, error) {
	params := url.Values{}
	params.Set("recvWindow", fmt.Sprint(40000))
	params.Set("timestamp", fmt.Sprint(time.Now().UnixMilli()))
	signature := api.HmacSha256(params.Encode(), os.Getenv("PRIVATE_API_KEY"))
	params.Set("signature", signature)

	res, err := b.api.Get(constant.BALANCE_URI+"?"+params.Encode(), false)
	if err != nil {
		return balance.Balance{}, err
	}

	data := []balance.Balance{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return balance.Balance{}, err
	}

	i := pie.FindFirstUsing(data, func(val balance.Balance) bool {
		return val.Asset == constant.USDT
	})
	if i == -1 {
		return balance.Balance{}, errors.New("usdt not found")
	}

	return data[i], nil
}
