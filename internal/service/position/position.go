package position

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/constant"
	entity "github.com/antongoncharik/crypto-knight-api/internal/entity/position"
	"github.com/antongoncharik/crypto-knight-api/pkg/api"
)

type PositionService struct {
	api *api.HTTPClient
}

func New(api *api.HTTPClient) *PositionService {
	return &PositionService{api: api}
}

func (p *PositionService) GetPositions() (entity.Positions, error) {
	params := url.Values{}
	params.Set("recvWindow", fmt.Sprint(40000))
	params.Set("timestamp", fmt.Sprint(time.Now().UnixMilli()))
	signature := api.HmacSha256(params.Encode(), os.Getenv("PRIVATE_API_KEY"))
	params.Set("signature", signature)

	var res []byte
	var err error
	res, err = p.api.Get(constant.POSITION_RISK_URI+"?"+params.Encode(), false)
	if err != nil {
		return nil, err
	}

	data := []entity.Position{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	positions := make(entity.Positions, len(data))
	for _, v := range data {
		positions[v.Symbol] = v
	}

	return positions, nil
}
