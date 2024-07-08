package marketdatasrests

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Lornzo/gobinance/commons"
	"github.com/shopspring/decimal"
)

func NewKLines(baseURL string) *KLines {
	var kLines *KLines = &KLines{
		Base: commons.NewBinanceRestfulBase(),
	}
	kLines.Base.SetBaseURL(baseURL)
	return kLines
}

type KLines struct {
	Base commons.RestfulBase
}

func (k *KLines) GetAPIPathes() []string {
	return []string{"fapi", "v1", "klines"}
}

func (k *KLines) GetWeights() int {

	var (
		limitStr string = k.Base.GetQuery("limit")
		limit    int64
		err      error
	)

	if limitStr == "" {
		return 2
	}

	if limit, err = strconv.ParseInt(limitStr, 10, 64); err != nil {
		return 0
	}

	if limit > 1000 {
		return 10
	}

	if limit > 500 {
		return 5
	}

	if limit > 100 {
		return 2
	}

	return 1

}

func (k *KLines) SetSymbol(symbol string) {
	k.Base.SetQuery("symbol", symbol)
}

func (k *KLines) SetInterval(interval string) {
	k.Base.SetQuery("interval", interval)
}

func (k *KLines) SetStartTimestamp(ts int64) {
	k.Base.SetQuery("startTime", fmt.Sprint(ts))
}

func (k *KLines) SetEndTimestamp(ts int64) {
	k.Base.SetQuery("endTime", fmt.Sprint(ts))
}

func (k *KLines) SetLimit(limit int64) {
	k.Base.SetQuery("limit", fmt.Sprint(limit))
}

func (k *KLines) DoRequest(ctx context.Context) (KLinesResponse, error) {

	var (
		resp   [][]interface{}
		err    error
		result KLinesResponse
	)

	k.initApiUrl()

	if _, err = k.Base.GET(ctx, &resp); err != nil {
		return KLinesResponse{}, err
	}

	for _, v := range resp {

		var candle KLine

		candle.OpenTimestamp = decimal.NewFromFloat(v[0].(float64)).BigInt().Int64()

		if openPrice, err := decimal.NewFromString(v[1].(string)); err != nil {
			return result, err
		} else {
			candle.OpenPrice = openPrice
		}

		if highPrice, err := decimal.NewFromString(v[2].(string)); err != nil {
			return result, err
		} else {
			candle.HighPrice = highPrice
		}

		if lowPrice, err := decimal.NewFromString(v[3].(string)); err != nil {
			return result, err
		} else {
			candle.LowPrice = lowPrice
		}

		if closePrice, err := decimal.NewFromString(v[4].(string)); err != nil {
			return result, err
		} else {
			candle.ClosePrice = closePrice
		}

		if volumn, err := decimal.NewFromString(v[5].(string)); err != nil {
			return result, err
		} else {
			candle.Volume = volumn
		}

		candle.CloseTimestamp = decimal.NewFromFloat(v[6].(float64)).BigInt().Int64()

		if turnOver, err := decimal.NewFromString(v[7].(string)); err != nil {
			return result, err
		} else {
			candle.Turnover = turnOver
		}

		candle.TurnoverNum = decimal.NewFromFloat(v[8].(float64)).BigInt().Int64()

		if activeBuyVolumn, err := decimal.NewFromString(v[9].(string)); err != nil {
			return result, err
		} else {
			candle.ActiveBuyTurn = activeBuyVolumn
		}

		if activeBuyTurnOver, err := decimal.NewFromString(v[10].(string)); err != nil {
			return result, err
		} else {
			candle.ActiveBuyVolume = activeBuyTurnOver
		}

		result = append(result, candle)

	}

	return result, nil
}

func (k *KLines) initApiUrl() {
	k.Base.SetPathes(k.GetAPIPathes()...)
}
