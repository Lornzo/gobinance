package marketdataswebsockets

import "github.com/shopspring/decimal"

type binanceKLine struct {
	binanceEvent
	Symbol    string `json:"s"`
	KLineData struct {
		StartMilliSecond int64           `json:"t"`
		CloseMilliSecond int64           `json:"T"`
		Symbol           string          `json:"s"`
		Interval         string          `json:"i"`
		FirstTradeID     int64           `json:"f"`
		LastTradeID      int64           `json:"L"`
		OpenPrice        decimal.Decimal `json:"o"`
		ClosePrice       decimal.Decimal `json:"c"`
		HighPrice        decimal.Decimal `json:"h"`
		LowPrice         decimal.Decimal `json:"l"`
		Volume           decimal.Decimal `json:"v"` // 成交量
		Number           decimal.Decimal `json:"n"` // 成交筆數
		Finished         bool            `json:"x"` // 是否結束
		Quata            decimal.Decimal `json:"q"` // 成交額
		ActiveBuyVolume  decimal.Decimal `json:"V"` // 主動買入成交量
		ActivebyQuata    decimal.Decimal `json:"Q"` // 主動買入成交額
		OtherB           decimal.Decimal `json:"B"` // 忽略
	} `json:"k"`
}

func (b binanceKLine) GetStartMilliSecond() int64 {
	return b.KLineData.StartMilliSecond
}

func (b binanceKLine) GetCloseMilliSecond() int64 {
	return b.KLineData.CloseMilliSecond
}

func (b binanceKLine) GetSymbol() string {
	return b.Symbol
}

func (b binanceKLine) GetInterval() string {
	return b.KLineData.Interval
}

func (b binanceKLine) GetOpenPrice() decimal.Decimal {
	return b.KLineData.OpenPrice
}

func (b binanceKLine) GetClosePrice() decimal.Decimal {
	return b.KLineData.ClosePrice
}

func (b binanceKLine) GetHighPrice() decimal.Decimal {
	return b.KLineData.HighPrice
}

func (b binanceKLine) GetLowPrice() decimal.Decimal {
	return b.KLineData.LowPrice
}

func (b binanceKLine) GetVolume() decimal.Decimal {
	return b.KLineData.Volume
}
