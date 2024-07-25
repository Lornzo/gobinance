package marketdataswebsockets

import "github.com/shopspring/decimal"

type AggTrade interface {
	GetEventName() string          // 事件类型
	GetEventTimestampMilli() int64 // 事件时间
	GetSymbol() string             // 交易对
	GetAddTradeID() int64          // 归集成交 ID
	GetPrice() decimal.Decimal     // 成交价格
	GetQuantity() decimal.Decimal  // 成交量
	GetFirstTradeID() int64        // 被归集的首个交易ID
	GetLastTradeID() int64         // 被归集的末次交易ID
	GetTimestampMilli() int64      // 成交时间
	IsMarketMaker() bool           // 买方是否是做市方。如true，则此次成交是一个主动卖出单，否则是一个主动买入单。
}

type aggTrade struct {
	EventName           string          `json:"e"`
	EventTimestampMilli int64           `json:"E"`
	Symbol              string          `json:"s"`
	AggTradeID          int64           `json:"a"`
	Price               decimal.Decimal `json:"p"`
	Quantity            decimal.Decimal `json:"q"`
	FirstTradeID        int64           `json:"f"`
	LastTradeID         int64           `json:"l"`
	TimestampMilli      int64           `json:"T"`
	MarketMaker         bool            `json:"m"`
}

func (a aggTrade) GetEventName() string {
	return a.EventName
}

func (a aggTrade) GetEventTimestampMilli() int64 {
	return a.EventTimestampMilli
}

func (a aggTrade) GetSymbol() string {
	return a.Symbol
}

func (a aggTrade) GetAddTradeID() int64 {
	return a.AggTradeID
}

func (a aggTrade) GetPrice() decimal.Decimal {
	return a.Price
}

func (a aggTrade) GetQuantity() decimal.Decimal {
	return a.Quantity
}

func (a aggTrade) GetFirstTradeID() int64 {
	return a.FirstTradeID
}

func (a aggTrade) GetLastTradeID() int64 {
	return a.LastTradeID
}

func (a aggTrade) GetTimestampMilli() int64 {
	return a.TimestampMilli
}

func (a aggTrade) IsMarketMaker() bool {
	return a.MarketMaker
}
