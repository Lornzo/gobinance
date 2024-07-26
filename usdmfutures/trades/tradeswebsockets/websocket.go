package tradeswebsockets

import (
	"context"

	"github.com/Lornzo/gobinance/binancewebsockets"
)

type Websocket interface {
	RunNewThread(ctx context.Context)
	PlaceOrder(ctx context.Context, account Account, order PlaceOrderForm) (Order, error)
	Close() error
}

func NewWebsocket(ctx context.Context, binanceWebsocketURL string) (Websocket, error) {

	var (
		ws  *tradeWebsocket = &tradeWebsocket{}
		err error
	)

	if ws.ws, err = binancewebsockets.NewWebsocket(binanceWebsocketURL); err != nil {
		return nil, err
	}

	return ws, nil

}
