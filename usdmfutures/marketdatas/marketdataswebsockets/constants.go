package marketdataswebsockets

const (
	REQUEST_METHOD_SUBSCRIBE   string = "SUBSCRIBE"
	REQUEST_METHOD_UNSUBSCRIBE string = "UNSUBSCRIBE"

	EVENT_TYPE_KLINE             string = "kline"
	EVENT_TYPE_AGG_TRADE         string = "aggTrade"
	EVENT_TYPE_MARK_PRICE_UPDATE string = "markPriceUpdate"
)
