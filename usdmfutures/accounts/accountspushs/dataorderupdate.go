package accountspushs

import "github.com/shopspring/decimal"

type DataOrder struct {
	// 特殊的自定义订单ID:
	// "autoclose-"开头的字符串: 系统强平订单
	// "adl_autoclose": ADL自动减仓订单
	// "settlement_autoclose-": 下架或交割的结算订单
	ClientID                  string          `json:"c"`   // 客户端自定订单ID
	Symbol                    string          `json:"s"`   // 交易对
	Side                      string          `json:"S"`   // 訂單方向
	OrderType                 string          `json:"o"`   // 订单类型
	TimeInForce               string          `json:"f"`   // 有效方式
	Quantity                  decimal.Decimal `json:"q"`   // 訂單原始數量
	Price                     decimal.Decimal `json:"p"`   // 訂單原始價格
	AveragePrice              decimal.Decimal `json:"ap"`  // 訂單平均價格
	TriggerPrice              decimal.Decimal `json:"sp"`  // 条件订单触发价格，对追踪止损单无效
	ExcuteType                string          `json:"x"`   // 本次事件的具体执行类型
	ExcuteStatus              string          `json:"X"`   // 订单的当前状态
	OrderID                   int64           `json:"i"`   // 订单ID
	LastAmount                decimal.Decimal `json:"l"`   // 订单末次成交量
	AccumulationAmount        decimal.Decimal `json:"z"`   // 訂單累計已成交量
	LastPrice                 decimal.Decimal `json:"L"`   // 订单末次成交价格
	FeeAssetType              string          `json:"N"`   // 手续费资产类型
	FeeAmount                 decimal.Decimal `json:"n"`   // 手续费数量
	TradeTimestamp            int64           `json:"T"`   // 成交时间
	TradeID                   int64           `json:"t"`   // 成交ID
	BuyOrderEquity            decimal.Decimal `json:"b"`   // 買單淨值
	SellOrderEquity           decimal.Decimal `json:"a"`   // 賣單淨值
	IsPendingOrderTransaction bool            `json:"m"`   // 该成交是作为挂单成交吗？
	IsReducePosition          bool            `json:"R"`   // 是否是只减仓单
	TriggerType               string          `json:"wt"`  // 触发价类型
	OriginOrderType           string          `json:"ot"`  // 原始订单类型
	PositionSide              string          `json:"ps"`  // 持倉方向
	IsTriggerCleanPosition    bool            `json:"cp"`  // 是否为触发平仓单; 仅在条件订单情况下会推送此字段
	TrackStopTriggerPrice     decimal.Decimal `json:"AP"`  // 追踪止损激活价格, 仅在追踪止损单时会推送此字段
	TrackStopTriggerRatio     decimal.Decimal `json:"cr"`  // 追踪止损回调比例, 仅在追踪止损单时会推送此字段
	IsProtectTrigger          bool            `json:"pP"`  // 是否开启条件单触发保护
	SI                        int64           `json:"si"`  // 未知，請忽略
	SS                        int64           `json:"ss"`  // 未知，請忽略
	RealizedProfit            decimal.Decimal `json:"rp"`  // 该交易实现盈亏
	SelfDealPreventionMode    string          `json:"V"`   // 自成交防止模式
	PriceMatchingMode         string          `json:"pm"`  // 价格匹配模式
	GTD                       int64           `json:"gtd"` // TIF为GTD的订单自动取消时间
}

type DataOrderUpdate struct {
	DataEvent
	Timestamp int64     `json:"T"` // 撮合时间
	Order     DataOrder `json:"o"`
}

func (d DataOrderUpdate) GetTimestamp() int64 {
	return d.Timestamp
}

func (d DataOrderUpdate) GetClientID() string {
	return d.Order.ClientID
}

func (d DataOrderUpdate) GetSymbol() string {
	return d.Order.Symbol
}

func (d DataOrderUpdate) GetOrder() string {
	return d.Order.Side
}

func (d DataOrderUpdate) GetOrderSide() string {
	return d.GetOrder()
}

func (d DataOrderUpdate) GetOrderType() string {
	return d.Order.OrderType
}

func (d DataOrderUpdate) GetTimeInForce() string {
	return d.Order.TimeInForce
}

func (d DataOrderUpdate) GetQuantity() decimal.Decimal {
	return d.Order.Quantity
}

func (d DataOrderUpdate) GetPrice() decimal.Decimal {
	return d.Order.Price
}

func (d DataOrderUpdate) GetAveragePrice() decimal.Decimal {
	return d.Order.AveragePrice
}

func (d DataOrderUpdate) GetTriggerPrice() decimal.Decimal {
	return d.Order.TriggerPrice
}

func (d DataOrderUpdate) GetExcuteType() string {
	return d.Order.ExcuteType
}

func (d DataOrderUpdate) GetExcuteStatus() string {
	return d.Order.ExcuteStatus
}

func (d DataOrderUpdate) GetOrderID() int64 {
	return d.Order.OrderID
}

func (d DataOrderUpdate) GetLastAmount() decimal.Decimal {
	return d.Order.LastAmount
}

func (d DataOrderUpdate) GetAccumulationAmount() decimal.Decimal {
	return d.Order.AccumulationAmount
}

func (d DataOrderUpdate) GetLastPrice() decimal.Decimal {
	return d.Order.LastPrice
}

func (d DataOrderUpdate) GetFeeAssetType() string {
	return d.Order.FeeAssetType
}

func (d DataOrderUpdate) GetFeeAmount() decimal.Decimal {
	return d.Order.FeeAmount
}

func (d DataOrderUpdate) GetTradeTimestamp() int64 {
	return d.Order.TradeTimestamp
}

func (d DataOrderUpdate) GetTradeID() int64 {
	return d.Order.TradeID
}

func (d DataOrderUpdate) GetBuyOrderEquity() decimal.Decimal {
	return d.Order.BuyOrderEquity
}

func (d DataOrderUpdate) GetSellOrderEquity() decimal.Decimal {
	return d.Order.SellOrderEquity
}

func (d DataOrderUpdate) GetIsPendingOrderTransaction() bool {
	return d.Order.IsPendingOrderTransaction
}

func (d DataOrderUpdate) GetIsReducePosition() bool {
	return d.Order.IsReducePosition
}

func (d DataOrderUpdate) GetTriggerType() string {
	return d.Order.TriggerType
}

func (d DataOrderUpdate) GetOriginOrderType() string {
	return d.Order.OriginOrderType
}

func (d DataOrderUpdate) GetPositionSide() string {
	return d.Order.PositionSide
}

func (d DataOrderUpdate) GetIsTriggerCleanPosition() bool {
	return d.Order.IsTriggerCleanPosition
}

func (d DataOrderUpdate) GetTrackStopTriggerPrice() decimal.Decimal {
	return d.Order.TrackStopTriggerPrice
}

func (d DataOrderUpdate) GetTrackStopTriggerRatio() decimal.Decimal {
	return d.Order.TrackStopTriggerRatio
}

func (d DataOrderUpdate) GetIsProtectTrigger() bool {
	return d.Order.IsProtectTrigger
}

func (d DataOrderUpdate) GetSI() int64 {
	return d.Order.SI
}

func (d DataOrderUpdate) GetSS() int64 {
	return d.Order.SS
}

func (d DataOrderUpdate) GetRealizedProfit() decimal.Decimal {
	return d.Order.RealizedProfit
}

func (d DataOrderUpdate) GetSelfDealPreventionMode() string {
	return d.Order.SelfDealPreventionMode
}

func (d DataOrderUpdate) GetPriceMatchingMode() string {
	return d.Order.PriceMatchingMode
}

func (d DataOrderUpdate) GetGTD() int64 {
	return d.Order.GTD
}

type OrderUpdate interface {
	GetEventName() string
	GetEventTs() int64
	GetTimestamp() int64
	GetClientID() string
	GetSymbol() string
	GetOrderSide() string
	GetOrderType() string
	GetTimeInForce() string
	GetQuantity() decimal.Decimal
	GetPrice() decimal.Decimal
	GetAveragePrice() decimal.Decimal
	GetTriggerPrice() decimal.Decimal
	GetExcuteType() string
	GetExcuteStatus() string
	GetOrderID() int64
	GetLastAmount() decimal.Decimal
	GetAccumulationAmount() decimal.Decimal
	GetLastPrice() decimal.Decimal
	GetFeeAssetType() string
	GetFeeAmount() decimal.Decimal
	GetTradeTimestamp() int64
	GetTradeID() int64
	GetBuyOrderEquity() decimal.Decimal
	GetSellOrderEquity() decimal.Decimal
	GetIsPendingOrderTransaction() bool
	GetIsReducePosition() bool
	GetTriggerType() string
	GetOriginOrderType() string
	GetPositionSide() string
	GetIsTriggerCleanPosition() bool
	GetTrackStopTriggerPrice() decimal.Decimal
	GetTrackStopTriggerRatio() decimal.Decimal
	GetIsProtectTrigger() bool
	GetSI() int64
	GetSS() int64
	GetRealizedProfit() decimal.Decimal
	GetSelfDealPreventionMode() string
	GetPriceMatchingMode() string
	GetGTD() int64
}
