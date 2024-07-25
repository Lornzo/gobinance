package marketdataswebsockets

import "github.com/shopspring/decimal"

type KLine interface {
	GetEventName() string
	GetEventTimestampMilli() int64
	GetOpenTimestampMilli() int64
	GetCloseTimestampMilli() int64
	GetSymbol() string
	GetInterval() string
	GetFirstTradeID() int64
	GetLastTradeID() int64
	GetOpenPrice() decimal.Decimal
	GetClosePrice() decimal.Decimal
	GetHighPrice() decimal.Decimal
	GetLowPrice() decimal.Decimal
	GetVolume() decimal.Decimal
	GetNumber() int64
	IsClosed() bool
	GetQuantity() decimal.Decimal
	GetActiveBuyingVolume() decimal.Decimal
	GetActiveBuyingQuantity() decimal.Decimal
}

type kLine struct {
	EventName           string `json:"e"`
	EventTimestampMilli int64  `json:"E"`
	Symbol              string `json:"s"`
	KLine               struct {
		OpenTimestampMilli   int64           `json:"t"`
		CloseTimestampMilli  int64           `json:"T"`
		Symbol               string          `json:"s"`
		Interval             string          `json:"i"`
		FirstTradeID         int64           `json:"f"`
		LastTradeID          int64           `json:"L"`
		OpenPrice            decimal.Decimal `json:"o"`
		ClosePrice           decimal.Decimal `json:"c"`
		HighPrice            decimal.Decimal `json:"h"`
		LowPrice             decimal.Decimal `json:"l"`
		Volume               decimal.Decimal `json:"v"` // 这根K线期间成交量
		Number               int64           `json:"n"` // 这根K线期间成交笔数
		Closed               bool            `json:"x"` // 这根K线是否完结(是否已经开始下一根K线)
		Quantity             decimal.Decimal `json:"q"` // 这根K线期间成交额
		ActiveBuyingVolume   decimal.Decimal `json:"V"` // 这根K线期间主动买入成交量
		ActiveBuyingQuantity decimal.Decimal `json:"Q"` // 这根K线期间主动买入成交额
		B                    decimal.Decimal `json:"B"` // ignore
	} `json:"k"`
}

func (k kLine) GetEventName() string {
	return k.EventName
}

func (k kLine) GetEventTimestampMilli() int64 {
	return k.EventTimestampMilli
}

func (k kLine) GetOpenTimestampMilli() int64 {
	return k.KLine.OpenTimestampMilli
}

func (k kLine) GetCloseTimestampMilli() int64 {
	return k.KLine.CloseTimestampMilli
}

func (k kLine) GetSymbol() string {
	return k.KLine.Symbol
}

func (k kLine) GetInterval() string {
	return k.KLine.Interval
}

func (k kLine) GetFirstTradeID() int64 {
	return k.KLine.FirstTradeID
}

func (k kLine) GetLastTradeID() int64 {
	return k.KLine.LastTradeID
}

func (k kLine) GetOpenPrice() decimal.Decimal {
	return k.KLine.OpenPrice
}

func (k kLine) GetClosePrice() decimal.Decimal {
	return k.KLine.ClosePrice
}

func (k kLine) GetHighPrice() decimal.Decimal {
	return k.KLine.HighPrice
}

func (k kLine) GetLowPrice() decimal.Decimal {
	return k.KLine.LowPrice
}

func (k kLine) GetVolume() decimal.Decimal {
	return k.KLine.Volume
}

func (k kLine) GetNumber() int64 {
	return k.KLine.Number
}

func (k kLine) IsClosed() bool {
	return k.KLine.Closed
}

func (k kLine) GetQuantity() decimal.Decimal {
	return k.KLine.Quantity
}

func (k kLine) GetActiveBuyingVolume() decimal.Decimal {
	return k.KLine.ActiveBuyingVolume
}

func (k kLine) GetActiveBuyingQuantity() decimal.Decimal {
	return k.KLine.ActiveBuyingQuantity
}
