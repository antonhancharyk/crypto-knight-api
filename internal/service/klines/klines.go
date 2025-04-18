package kline

import (
	"encoding/json"
	"net/url"

	"github.com/antongoncharik/crypto-knight-api/internal/constant"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
)

type KlinerService struct {
	api *api.HTTPClient
}

func New(api *api.HTTPClient) *KlinerService {
	return &KlinerService{api: api}
}

func (k *KlinerService) Get(sbl string) ([][]any, error) {
	params := url.Values{}
	params.Set("symbol", sbl)
	params.Set("interval", constant.INTERVAL_KLINES)
	params.Set("limit", constant.QUANTITY_KLINES)

	res, err := k.api.Get(constant.KLINES_URI+"?"+params.Encode(), false)
	if err != nil {
		return [][]any{}, err
	}

	data := [][]any{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return [][]any{}, err
	}

	return data, nil
}
