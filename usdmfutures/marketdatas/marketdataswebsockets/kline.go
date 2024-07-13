package marketdataswebsockets

import "github.com/shopspring/decimal"

type KLine interface {
	GetStartMilliSecond() int64
	GetCloseMilliSecond() int64
	GetSymbol() string
	GetInterval() string
	GetOpenPrice() decimal.Decimal
	GetClosePrice() decimal.Decimal
	GetHighPrice() decimal.Decimal
	GetLowPrice() decimal.Decimal
	GetVolume() decimal.Decimal
}
