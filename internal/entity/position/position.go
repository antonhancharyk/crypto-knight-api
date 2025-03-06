package entity

type Position struct {
	Symbol                 string `json:"symbol"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	EntryPrice             string `json:"entryPrice"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	MarkPrice              string `json:"markPrice"`
	UnRealizedProfit       string `json:"unRealizedProfit"`
	LiquidationPrice       string `json:"liquidationPrice"`
	IsolatedMargin         string `json:"isolatedMargin"`
	Notional               string `json:"notional"`
	MarginAsset            string `json:"marginAsset"`
	IsolatedWallet         string `json:"isolatedWallet"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Adl                    int64  `json:"adl"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	UpdateTime             int64  `json:"updateTime"`
}

type Positions map[string]Position
