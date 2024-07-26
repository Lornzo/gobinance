package tradeswebsockets

import (
	"context"
	"encoding/json"

	"github.com/Lornzo/gobinance/binancewebsockets"
)

type requestBuilder interface {
	buildRequestBody() (map[string]interface{}, error)
}

type tradeWebsocket struct {
	ws binancewebsockets.Websocket
}

func (t *tradeWebsocket) RunNewThread(ctx context.Context) {
	t.ws.RunNewThread(ctx)
}

func (t *tradeWebsocket) Run(ctx context.Context) {
	t.ws.Run(ctx)
}

func (t *tradeWebsocket) PlaceOrder(ctx context.Context, account Account, order PlaceOrderForm) (Order, error) {
	return t.requestOrder(ctx, trade_method_place, newFormBuilder(account, placeOrderFormBuilder{order: order}))
}

func (t *tradeWebsocket) ModifyOrder(ctx context.Context, account Account, order ModifyOrderForm) (Order, error) {
	return t.requestOrder(ctx, trade_method_modify, newFormBuilder(account, modifyOrderFormBuilder{order: order}))
}

func (t *tradeWebsocket) CancelOrder(ctx context.Context, account Account, order CancelOrderForm) (Order, error) {
	return t.requestOrder(ctx, trade_method_cancel, newFormBuilder(account, cancelOrderFormBuilder{order: order}))
}

func (t *tradeWebsocket) StatusOrder(ctx context.Context, account Account, order StatusOrderForm) (Order, error) {
	return t.requestOrder(ctx, trade_method_status, newFormBuilder(account, statusOrderFormBuilder{order: order}))
}

func (t *tradeWebsocket) AccountPosition(ctx context.Context, account Account, position PositionForm) (Positions, error) {
	return t.requestPosition(ctx, trade_method_account_position, newFormBuilder(account, positionFormBuilder{form: position}))
}

func (t *tradeWebsocket) AccountPositionV2(ctx context.Context, account Account, position PositionForm) (PositionsV2, error) {
	return t.requestPositionV2(ctx, trade_method_account_position_v2, newFormBuilder(account, positionFormBuilder{form: position}))
}

func (t *tradeWebsocket) requestOrder(ctx context.Context, method string, requestBuilder requestBuilder) (Order, error) {

	var (
		jsonBytes     []byte
		orderResponse orderResponse
		err           error
	)

	if jsonBytes, err = t.requestBytes(ctx, method, requestBuilder); err != nil {
		return Order{}, err
	}

	if err = json.Unmarshal(jsonBytes, &orderResponse); err != nil {
		return Order{}, err
	}

	return orderResponse.Result, nil
}

func (t *tradeWebsocket) requestPosition(ctx context.Context, method string, requestBuilder requestBuilder) (Positions, error) {

	var (
		jsonBytes        []byte
		err              error
		positionResponse struct {
			ID     string    `json:"id"`
			Status int       `json:"status"`
			Result Positions `json:"result"`
		}
	)

	if jsonBytes, err = t.requestBytes(ctx, method, requestBuilder); err != nil {
		return Positions{}, err
	}

	if err = json.Unmarshal(jsonBytes, &positionResponse); err != nil {
		return Positions{}, err
	}

	return positionResponse.Result, nil

}

func (t *tradeWebsocket) requestPositionV2(ctx context.Context, method string, requestBuilder requestBuilder) (PositionsV2, error) {
	var (
		jsonBytes        []byte
		err              error
		positionResponse struct {
			ID     string      `json:"id"`
			Status int         `json:"status"`
			Result PositionsV2 `json:"result"`
		}
	)

	if jsonBytes, err = t.requestBytes(ctx, method, requestBuilder); err != nil {
		return PositionsV2{}, err
	}

	if err = json.Unmarshal(jsonBytes, &positionResponse); err != nil {
		return PositionsV2{}, err
	}

	return positionResponse.Result, nil
}

func (t *tradeWebsocket) requestBytes(ctx context.Context, method string, requestBuilder requestBuilder) ([]byte, error) {

	var (
		response binancewebsockets.Response
		params   map[string]interface{}
		err      error
	)

	if params, err = requestBuilder.buildRequestBody(); err != nil {
		return nil, err
	}

	if response, err = t.ws.MakeRequestByUUIDIndex(ctx, binancewebsockets.BinanceRequest{
		Method: method,
		Params: params,
	}); err != nil {
		return nil, err
	}

	if err = response.Err; err != nil {
		return nil, err
	}

	return response.Msg, nil

}

func (t *tradeWebsocket) Close() error {
	return t.ws.Close()
}
