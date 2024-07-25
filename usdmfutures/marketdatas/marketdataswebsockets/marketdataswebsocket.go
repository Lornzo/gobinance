package marketdataswebsockets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Lornzo/gobinance/binancewebsockets"
	"github.com/Lornzo/gobinance/threadsafetypes"
	"github.com/Lornzo/gobinance/usdmfutures/usdmfutureswebsockettypes"
)

type marketDatasWebsocket struct {
	ws                   binancewebsockets.Websocket
	streamIDs            threadsafetypes.MapStringInterface
	errorSubscribers     errorSubscribers
	kLineSubscribers     kLineSubscribersStream
	aggTradeSubscribers  aggTradeSubscribersStream
	markPriceSubscribers markPriceSubscribersStream
}

func (m *marketDatasWebsocket) Run(ctx context.Context) {
	m.ws.Run(ctx, m.runHandler)
}

func (m *marketDatasWebsocket) RunNewThread(ctx context.Context) {
	m.ws.RunNewThread(ctx, m.runHandler)
}

func (m *marketDatasWebsocket) SubscribeError(subscriber usdmfutureswebsockettypes.ErrorSubscriber) error {
	return m.errorSubscribers.Subscribe(subscriber)
}

func (m *marketDatasWebsocket) UnSubscribeError(subscriber usdmfutureswebsockettypes.ErrorSubscriber) error {
	return m.errorSubscribers.UnSubscribe(subscriber)
}

func (m *marketDatasWebsocket) SubscribeKLine(ctx context.Context, subscriber KLineSubscriber) error {

	var (
		streamName      string = m.kLineSubscribers.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())
		streamNameExist bool   = m.kLineSubscribers.existStreamName(streamName)
		err             error
	)

	if err = m.kLineSubscribers.subscribe(subscriber); err != nil {
		return err
	}

	if streamNameExist {
		return nil
	}

	return m.subscribe(ctx, streamName)

}

func (m *marketDatasWebsocket) UnSubscribeKLine(ctx context.Context, subscriber KLineSubscriber) error {

	var (
		streamName string = m.kLineSubscribers.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())
		err        error
	)

	if err = m.kLineSubscribers.unsubscribe(subscriber); err != nil {
		return err
	}

	if m.kLineSubscribers.existStreamName(streamName) {
		return nil
	}

	return m.unsubscribe(ctx, streamName)

}

func (m *marketDatasWebsocket) SubscribeAggTrade(ctx context.Context, subscriber AggTradeSubscriber) error {

	var (
		streamName      string = m.aggTradeSubscribers.getStreamName(subscriber.GetSymbol())
		streamNameExist bool   = m.aggTradeSubscribers.existStreamName(streamName)
		err             error
	)

	if err = m.aggTradeSubscribers.subscribe(subscriber); err != nil {
		return err
	}

	if streamNameExist {
		return nil
	}

	return m.subscribe(ctx, streamName)
}

func (m *marketDatasWebsocket) UnSubscribeAggTrade(ctx context.Context, subscriber AggTradeSubscriber) error {

	var (
		streamName string = m.aggTradeSubscribers.getStreamName(subscriber.GetSymbol())
		err        error
	)

	if err = m.aggTradeSubscribers.unsubscribe(subscriber); err != nil {
		return err
	}

	if m.aggTradeSubscribers.existStreamName(streamName) {
		return nil
	}

	return m.unsubscribe(ctx, streamName)
}

func (m *marketDatasWebsocket) SubscribeMarkPrice(ctx context.Context, subscriber MarkPriceSubscriber) error {

	var (
		streamName      string = m.markPriceSubscribers.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())
		streamNameExist bool   = m.markPriceSubscribers.existStreamName(streamName)
		err             error
	)

	if err = m.markPriceSubscribers.subscribe(subscriber); err != nil {
		return err
	}

	if streamNameExist {
		return nil
	}

	return m.subscribe(ctx, streamName)

}

func (m *marketDatasWebsocket) UnSubscribeMarkPrice(ctx context.Context, subscriber MarkPriceSubscriber) error {

	var (
		streamName string = m.markPriceSubscribers.getStreamName(subscriber.GetSymbol(), subscriber.GetInterval())
		err        error
	)

	if err = m.markPriceSubscribers.unsubscribe(subscriber); err != nil {
		return err
	}

	if m.markPriceSubscribers.existStreamName(streamName) {
		return nil
	}

	return m.unsubscribe(ctx, streamName)

}

func (m *marketDatasWebsocket) subscribe(ctx context.Context, streamName string) error {

	var (
		response binancewebsockets.Response
		err      error
	)

	if response, err = m.ws.MakeRequestByIntIndex(ctx, binancewebsockets.BinanceRequest{
		Method: REQUEST_METHOD_SUBSCRIBE,
		Params: []string{streamName},
	}); err != nil {
		return err
	}

	m.streamIDs.Set(streamName, response.RequestID)

	return nil
}

func (m *marketDatasWebsocket) unsubscribe(ctx context.Context, streamName string) error {

	var (
		requestID interface{} = m.streamIDs.Get(streamName)
		err       error
	)

	if requestID == nil {
		return nil
	}

	if _, err = m.ws.MakeRequest(ctx, requestID, binancewebsockets.BinanceRequest{
		Method: REQUEST_METHOD_UNSUBSCRIBE,
		Params: []string{streamName},
	}); err != nil {
		return err
	}

	m.streamIDs.Delete(streamName)

	return nil

}

func (m *marketDatasWebsocket) runHandler(msgType int, msg []byte, msgError error) {

	var (
		resp struct {
			StreamName string                 `json:"stream"`
			Data       map[string]interface{} `json:"data"`
		}
		data []byte
		err  error
	)

	if msgType == -1 {
		m.disconnectHandler(msgType, msgError)
		return
	}

	if msgError != nil {
		m.errorSubscribers.UpdateError(msgError)
		return
	}

	if err = json.Unmarshal(msg, &resp); err != nil {
		m.errorSubscribers.UpdateError(err)
		return
	}

	if data, err = json.Marshal(resp.Data); err != nil {
		m.errorSubscribers.UpdateError(err)
		return

	}

	if eventName, nameExist := resp.Data["e"]; nameExist {
		m.eventNameHandler(resp.StreamName, fmt.Sprint(eventName), data)
	}

}

func (m *marketDatasWebsocket) disconnectHandler(msgType int, err error) {}

func (m *marketDatasWebsocket) eventNameHandler(streamName string, eventName string, data []byte) {
	switch eventName {
	case EVENT_TYPE_KLINE:
		m.kLineHandler(data)
		return
	case EVENT_TYPE_AGG_TRADE:
		m.aggTradeHandler(data)
		return
	case EVENT_TYPE_MARK_PRICE_UPDATE:
		m.markPriceHandler(streamName, data)
		return
	}
}

func (m *marketDatasWebsocket) kLineHandler(kLineData []byte) {

	var (
		kLine kLine
		err   error
	)

	if err = json.Unmarshal(kLineData, &kLine); err != nil {
		m.errorSubscribers.UpdateError(err)
	}

	m.kLineSubscribers.updateKLine(kLine, nil)

}

func (m *marketDatasWebsocket) aggTradeHandler(aggTradeData []byte) {

	var (
		aggTrade aggTrade
		err      error
	)

	if err = json.Unmarshal(aggTradeData, &aggTrade); err != nil {
		m.errorSubscribers.UpdateError(err)
	}

	m.aggTradeSubscribers.updateAggTrade(aggTrade, nil)

}

func (m *marketDatasWebsocket) markPriceHandler(streamName string, markPriceData []byte) {

	var (
		markPrice markPrice
		err       error
	)

	if err = json.Unmarshal(markPriceData, &markPrice); err != nil {
		m.errorSubscribers.UpdateError(err)
	}

	m.markPriceSubscribers.updateMarkPrice(streamName, markPrice, nil)

}

func (m *marketDatasWebsocket) Close() error {
	return m.ws.Close()
}
