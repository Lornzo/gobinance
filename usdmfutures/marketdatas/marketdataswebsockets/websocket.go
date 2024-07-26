package marketdataswebsockets

import (
	"context"
	"fmt"

	"github.com/Lornzo/gobinance/binancewebsockets"
	"github.com/Lornzo/gobinance/usdmfutures/usdmfutureswebsockettypes"
)

type Websocket interface {
	Close() error
	RunNewThread(ctx context.Context)
	Run(ctx context.Context)
	SubscribeKLine(ctx context.Context, subscriber KLineSubscriber) error
	UnSubscribeKLine(ctx context.Context, subscriber KLineSubscriber) error
	SubscribeMarkPrice(ctx context.Context, subscriber MarkPriceSubscriber) error
	UnSubscribeMarkPrice(ctx context.Context, subscriber MarkPriceSubscriber) error
}

func NewWebsocket(ctx context.Context, binanceWebsocketURL string) (Websocket, error) {

	var (
		ws  *marketDatasWebsocket = &marketDatasWebsocket{}
		err error
	)

	ws.errorSubscribers = &usdmfutureswebsockettypes.ErrorSubscribers{}
	if ws.ws, err = binancewebsockets.NewWebsocket(fmt.Sprint(binanceWebsocketURL, "/stream")); err != nil {
		return nil, err
	}

	return ws, nil

}
