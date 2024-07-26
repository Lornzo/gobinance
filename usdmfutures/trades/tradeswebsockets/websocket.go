package tradeswebsockets

import (
	"context"

	"github.com/Lornzo/gobinance/binancewebsockets"
)

type Websocket interface {
	RunNewThread(ctx context.Context)
	Run(ctx context.Context)
	Close() error
	PlaceOrder(ctx context.Context, account Account, order PlaceOrderForm) (Order, error)
	ModifyOrder(ctx context.Context, account Account, order ModifyOrderForm) (Order, error)
	CancelOrder(ctx context.Context, account Account, order CancelOrderForm) (Order, error)
	StatusOrder(ctx context.Context, account Account, order StatusOrderForm) (Order, error)
	AccountPosition(ctx context.Context, account Account, position PositionForm) (Positions, error)
	AccountPositionV2(ctx context.Context, account Account, position PositionForm) (PositionsV2, error)
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
