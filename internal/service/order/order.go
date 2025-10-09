package order

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/constant"
	"github.com/antongoncharik/crypto-knight-api/internal/entity/order"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
)

type OrderService struct {
	api *api.HTTPClient
}

func New(api *api.HTTPClient) *OrderService {
	return &OrderService{api: api}
}

func (p *OrderService) GetOpenOrders() ([]order.Order, error) {
	params := url.Values{}
	params.Set("recvWindow", fmt.Sprint(40000))
	params.Set("timestamp", fmt.Sprint(time.Now().UnixMilli()))
	signature := api.HmacSha256(params.Encode(), os.Getenv("PRIVATE_API_KEY"))
	params.Set("signature", signature)

	var res []byte
	var err error
	res, err = p.api.Get(constant.OPEN_ORDERS_URI+"?"+params.Encode(), false)
	if err != nil {
		return nil, err
	}

	data := []order.Order{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
