package accountswebsockets

import "github.com/shopspring/decimal"

type OrderTradeUpdate interface {
	GetEventName() string                      // 事件類型
	GetEventTimestampMilli() int64             // 事件時間
	GetTimestampMilli() int64                  // 撮合時間
	GetClientID() string                       // 客户端自定订单ID
	GetSymbol() string                         // 交易对
	GetSide() string                           // 訂單方向
	GetOrderType() string                      // 訂單類型
	GetTimeInForce() string                    // 有效方式
	GetQuantity() decimal.Decimal              // 訂單原始數量
	GetPrice() decimal.Decimal                 // 訂單原始價格
	GetAveragePrice() decimal.Decimal          // 訂單平均價格
	GetTriggerPrice() decimal.Decimal          // 条件订单触发价格，对追踪止损单无效
	GetExcuteType() string                     // 本次事件的具体执行类型
	GetExcuteStatus() string                   // 订单的当前状态
	GetOrderID() int64                         // 订单ID
	GetLastAmount() decimal.Decimal            // 订单末次成交量
	GetAccumulationAmount() decimal.Decimal    // 訂單累計已成交量
	GetLastPrice() decimal.Decimal             // 订单末次成交价格
	GetFeeAssetType() string                   // 手续费资产类型
	GetFeeAmount() decimal.Decimal             // 手续费数量
	GetTradeTimestampMilli() int64             // 成交时间
	GetTradeID() int64                         // 成交ID
	GetBuyOrderEquity() decimal.Decimal        // 買單淨值
	GetSellOrderEquity() decimal.Decimal       // 賣單淨值
	GetIsPendingOrderTransaction() bool        // 该成交是作为挂单成交吗？
	GetIsReducePosition() bool                 // 是否是只减仓单
	GetTriggerType() string                    // 触发价类型
	GetOriginOrderType() string                // 原始订单类型
	GetPositionSide() string                   // 持倉方向
	GetIsTriggerCleanPosition() bool           // 是否为触发平仓单; 仅在条件订单情况下会推送此字段
	GetTrackStopTriggerPrice() decimal.Decimal // 追踪止损激活价格, 仅在追踪止损单时会推送此字段
	GetTrackStopTriggerRatio() decimal.Decimal // 追踪止损回调比例, 仅在追踪止损单时会推送此字段
	GetIsProtectTrigger() bool                 // 是否开启条件单触发保护
	GetSI() int64                              // 未知，請忽略
	GetSS() int64                              // 未知，請忽略
	GetRealizedProfit() decimal.Decimal        // 该交易实现盈亏
	GetSelfDealPreventionMode() string         // 自成交防止模式
	GetPriceMatchingMode() string              // 价格匹配模式
	GetGTD() int64                             // TIF为GTD的订单自动取消时间
}

type orderTradeUpdate struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	TimestampMilli      int64  `json:"T"`
	Order               struct {
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
	} `json:"o"`
}

func (o orderTradeUpdate) GetEventName() string {
	return o.EventName
}

func (o orderTradeUpdate) GetEventTimestampMilli() int64 {
	return o.EventTimestampMilli
}

func (o orderTradeUpdate) GetTimestampMilli() int64 {
	return o.TimestampMilli
}

func (o orderTradeUpdate) GetClientID() string {
	return o.Order.ClientID
}

func (o orderTradeUpdate) GetSymbol() string {
	return o.Order.Symbol
}

func (o orderTradeUpdate) GetSide() string {
	return o.Order.Side
}

func (o orderTradeUpdate) GetOrderType() string {
	return o.Order.OrderType
}

func (o orderTradeUpdate) GetTimeInForce() string {
	return o.Order.TimeInForce
}

func (o orderTradeUpdate) GetQuantity() decimal.Decimal {
	return o.Order.Quantity
}

func (o orderTradeUpdate) GetPrice() decimal.Decimal {
	return o.Order.Price
}

func (o orderTradeUpdate) GetAveragePrice() decimal.Decimal {
	return o.Order.AveragePrice
}

func (o orderTradeUpdate) GetTriggerPrice() decimal.Decimal {
	return o.Order.TriggerPrice
}

func (o orderTradeUpdate) GetExcuteType() string {
	return o.Order.ExcuteType
}

func (o orderTradeUpdate) GetExcuteStatus() string {
	return o.Order.ExcuteStatus
}

func (o orderTradeUpdate) GetOrderID() int64 {
	return o.Order.OrderID
}

func (o orderTradeUpdate) GetLastAmount() decimal.Decimal {
	return o.Order.LastAmount
}

func (o orderTradeUpdate) GetAccumulationAmount() decimal.Decimal {
	return o.Order.AccumulationAmount
}

func (o orderTradeUpdate) GetLastPrice() decimal.Decimal {
	return o.Order.LastPrice
}

func (o orderTradeUpdate) GetFeeAssetType() string {
	return o.Order.FeeAssetType
}

func (o orderTradeUpdate) GetFeeAmount() decimal.Decimal {
	return o.Order.FeeAmount
}

func (o orderTradeUpdate) GetTradeTimestampMilli() int64 {
	return o.Order.TradeTimestamp
}

func (o orderTradeUpdate) GetTradeID() int64 {
	return o.Order.TradeID
}

func (o orderTradeUpdate) GetBuyOrderEquity() decimal.Decimal {
	return o.Order.BuyOrderEquity
}

func (o orderTradeUpdate) GetSellOrderEquity() decimal.Decimal {
	return o.Order.SellOrderEquity
}

func (o orderTradeUpdate) GetIsPendingOrderTransaction() bool {
	return o.Order.IsPendingOrderTransaction
}

func (o orderTradeUpdate) GetIsReducePosition() bool {
	return o.Order.IsReducePosition
}

func (o orderTradeUpdate) GetTriggerType() string {
	return o.Order.TriggerType
}

func (o orderTradeUpdate) GetOriginOrderType() string {
	return o.Order.OriginOrderType
}

func (o orderTradeUpdate) GetPositionSide() string {
	return o.Order.PositionSide
}

func (o orderTradeUpdate) GetIsTriggerCleanPosition() bool {
	return o.Order.IsTriggerCleanPosition
}

func (o orderTradeUpdate) GetTrackStopTriggerPrice() decimal.Decimal {
	return o.Order.TrackStopTriggerPrice
}

func (o orderTradeUpdate) GetTrackStopTriggerRatio() decimal.Decimal {
	return o.Order.TrackStopTriggerRatio
}

func (o orderTradeUpdate) GetIsProtectTrigger() bool {
	return o.Order.IsProtectTrigger
}

func (o orderTradeUpdate) GetSI() int64 {
	return o.Order.SI
}

func (o orderTradeUpdate) GetSS() int64 {
	return o.Order.SS
}

func (o orderTradeUpdate) GetRealizedProfit() decimal.Decimal {
	return o.Order.RealizedProfit
}

func (o orderTradeUpdate) GetSelfDealPreventionMode() string {
	return o.Order.SelfDealPreventionMode
}

func (o orderTradeUpdate) GetPriceMatchingMode() string {
	return o.Order.PriceMatchingMode
}

func (o orderTradeUpdate) GetGTD() int64 {
	return o.Order.GTD
}
