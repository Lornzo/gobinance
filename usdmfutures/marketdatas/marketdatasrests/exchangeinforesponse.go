package marketdatasrests

import "github.com/shopspring/decimal"

type ExchangeInfoResponseSymbolFilter struct{}

type ExchangeInfoResponseSymbolFilters []ExchangeInfoResponseSymbolFilter

type ExchangeInfoResponseSymbol struct {
	Symbol                string                            `json:"symbol"`                // 交易對
	Pair                  string                            `json:"pair"`                  // 標的交易對
	ContractType          string                            `json:"contractType"`          // 合约类型
	DeliveryDate          int64                             `json:"deliveryDate"`          // 交割日期
	OnboardDate           int64                             `json:"onboardDate"`           // 上线日期
	SymbolStatus          string                            `json:"status"`                // 交易对状态
	MaintMarginPercent    decimal.Decimal                   `json:"maintMarginPercent"`    // 未知，可忽略
	RequiredMarginPercent decimal.Decimal                   `json:"requiredMarginPercent"` // 未知，可忽略
	BaseAsset             string                            `json:"baseAsset"`             // 标的资产
	QuoteAsset            string                            `json:"quoteAsset"`            // 报价资产
	MarginAsset           string                            `json:"marginAsset"`           // 保证金资产
	PricePrecision        int64                             `json:"pricePrecision"`        // 价格小数点位数(仅作为系统精度使用，注意同tickSize 区分）
	QuantityPrecision     int64                             `json:"quantityPrecision"`     // 数量小数点位数(仅作为系统精度使用，注意同stepSize 区分）
	BaseAssetPrecision    int64                             `json:"baseAssetPrecision"`    // 标的资产精度
	QuotePrecision        int64                             `json:"quotePrecision"`        // 报价资产精度
	UnderlyingType        string                            `json:"underlyingType"`
	UnderlyingSubType     []string                          `json:"underlyingSubType"`
	SettlePlan            int64                             `json:"settlePlan"`
	TriggerProtect        decimal.Decimal                   `json:"triggerProtect"` // 开启"priceProtect"的条件订单的触发阈值
	Filters               ExchangeInfoResponseSymbolFilters `json:"filters"`
	OrderType             []string                          `json:"OrderType"`       // 訂單類型
	TimeInForce           []string                          `json:"timeInForce"`     // 有效方式
	LiquidationFee        decimal.Decimal                   `json:"liquidationFee"`  // 强平费率
	MarketTakeBound       decimal.Decimal                   `json:"marketTakeBound"` // 市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
}

type ExchangeinfoResponseSymbols []ExchangeInfoResponseSymbol

type ExchangeInfoResponseAsset struct {
	Asset             string      `json:"asset"`
	MarginAvailable   bool        `json:"marginAvailable"`   // 是否可用作保证金
	AutoAssetExchange interface{} `json:"autoAssetExchange"` // 保证金资产自动兑换阈值
}

type ExchangeInfoResponseAssets []ExchangeInfoResponseAsset

type ExchangeInfoResponseRateLimit struct {
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type ExchangeInfoResponseRateLimits []ExchangeInfoResponseRateLimit

type ExchangeInfoResponse struct {
	ExchangeFilters []interface{}                  `json:"exchangeFilters"`
	RateLimits      ExchangeInfoResponseRateLimits `json:"rateLimits"`
	ServerTimestamp int64                          `json:"serverTime"`
	Assets          ExchangeInfoResponseAssets     `json:"assets"`
	Symbols         ExchangeinfoResponseSymbols    `json:"symbols"`
	Timezone        string                         `json:"timezone"`
}
