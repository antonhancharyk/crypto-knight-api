package constant

const (
	BASE_BINANCE_URI    = "https://fapi.binance.com"
	KLINES_URI          = BASE_BINANCE_URI + "/fapi/v1/klines"
	EXCHANGE_INFO_URI   = BASE_BINANCE_URI + "/fapi/v1/exchangeInfo"
	PRICE_URI           = BASE_BINANCE_URI + "/fapi/v1/ticker/price"
	ORDER_URI           = BASE_BINANCE_URI + "/fapi/v1/order"
	OPEN_ORDERS_URI     = BASE_BINANCE_URI + "/fapi/v1/openOrders"
	BALANCE_URI         = BASE_BINANCE_URI + "/fapi/v2/balance"
	SERVER_TIME_URI     = BASE_BINANCE_URI + "/fapi/v1/time"
	POSITION_RISK_URI   = BASE_BINANCE_URI + "/fapi/v3/positionRisk"
	USER_TRADE_URI      = BASE_BINANCE_URI + "/fapi/v1/userTrades"
	BASE_API_URI        = "https://api.crypto-knight.online"
	STATUS_BOT_URI      = BASE_API_URI + "/common/status"
	TRACKS_BOT_URI      = BASE_API_URI + "/tracks"
	TRACKS_BULK_BOT_URI = BASE_API_URI + "/tracks/bulk"
	ENTRIES_BOT_URI     = BASE_API_URI + "/entries"
	WSS_BINANCE         = "wss://fstream.binance.com/ws/"
	USDT                = "USDT"
)
