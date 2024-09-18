package marketdatasreststest

import (
	"context"
	"testing"

	"github.com/Lornzo/gobinance/gobinanceenums"
	"github.com/Lornzo/gobinance/usdmfutures/marketdatas/marketdatasrests"
)

func TestKLinesDoRequest(t *testing.T) {
	var (
		api *marketdatasrests.KLines = marketdatasrests.NewKLines("https://fapi.binance.com")
		ctx context.Context          = context.Background()
	)

	api.SetSymbol("ethusdt")
	api.SetInterval(gobinanceenums.INTERVAL_FIVE_MINUTES)
	api.SetLimit(100)

	if data, err := api.DoRequest(ctx); err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}

}
