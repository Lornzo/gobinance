package accountspushs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Lornzo/gobinance/commons"
	"github.com/gorilla/websocket"
)

func NewListener(baseURL string, listenKey string) *Listener {
	var listner *Listener = &Listener{
		Base:        commons.NewBinanceWebsocketBase(),
		subscribers: make(map[string]ListenerSubscriber),
	}
	listner.Base.SetBaseURL(baseURL)
	listner.Base.SetPathes("ws", listenKey)
	return listner
}

type Listener struct {
	Base            commons.WebsocketBase
	wsConn          *websocket.Conn
	isRunning       bool
	subscribers     map[string]ListenerSubscriber
	subscribersLock sync.RWMutex
	isDebug         bool
	isDebugLock     sync.RWMutex
}

func (l *Listener) Debug(isDebug bool) {
	l.isDebugLock.Lock()
	defer l.isDebugLock.Unlock()
	l.isDebug = isDebug
}

func (l *Listener) IsDebug() bool {
	l.isDebugLock.RLock()
	defer l.isDebugLock.RUnlock()
	return l.isDebug
}

func (l *Listener) Start(ctx context.Context) error {

	if l.isRunning {
		return errors.New("listener is already running")
	}

	var (
		wsConnResp *http.Response
		err        error
	)

	if l.wsConn, wsConnResp, err = websocket.DefaultDialer.Dial(l.Base.GetWebsocketURL(), nil); err != nil {
		return err
	}

	if wsConnResp.StatusCode != http.StatusSwitchingProtocols {
		return fmt.Errorf("wsConn faild, status code: %d", wsConnResp.StatusCode)
	}

	go func() {
		l.listen()
	}()

	return nil

}

func (l *Listener) listen() {

	for {
		var msgType int
		var msg []byte
		var msgErr error
		var typeIdentifier map[string]interface{}
		var eventName string

		if msgType, msg, msgErr = l.wsConn.ReadMessage(); msgErr != nil {
			l.errorHandler(msgErr)
			continue
		}

		if msgType == websocket.CloseMessage || msgType == -1 {
			l.closeHandler()
			break
		}

		if l.IsDebug() {
			log.Print("accountspushs Listener accept message : ", string(msg), "\n")
		}

		if msgErr = json.Unmarshal(msg, &typeIdentifier); msgErr != nil {
			l.errorHandler(msgErr)
			continue
		}

		if v, exist := typeIdentifier["e"]; !exist {
			l.errorHandler(errors.New("event name not found"))
			continue
		} else {
			eventName = v.(string)
		}

		switch eventName {
		case EVENT_TYPE_LISTEN_KEY_EXPIRED:
			l.listenKeyExpiredHandler(msg)
		case EVENT_TYPE_ACCOUNT_UPDATE:
			l.accountUpdateHandler(msg)
		case EVENT_TYPE_MARGIN_CALL:
			l.marginCallHandler(msg)
		case EVENT_TYPE_ORDER_TRADE_UPDATE:
			l.orderTradeUpdateHandler(msg)
		case EVENT_TYPE_ACCOUNT_CONFIG_UPDATE:
			l.accountConfigUpdateHandler(msg)
		case EVENT_TYPE_STRATEGY_UPDATE:
			l.strategyUpdateHandler(msg)
		case EVENT_TYPE_GRID_UPDATE:
			l.gridUpdateHandler(msg)
		case EVENT_TYPE_CONDITIONAL_ORDER_TRIGGER_REJECT:
			l.conditionalOrderTriggerRejectHandler(msg)
		default:
			l.undefinedTypeHandler(msg)
		}

	}

}

func (l *Listener) Subscribe(subscriber ListenerSubscriber) error {
	l.subscribersLock.Lock()
	defer l.subscribersLock.Unlock()

	if _, exist := l.subscribers[subscriber.GetSubscriberID()]; exist {
		return fmt.Errorf("subscriber has already exist : %s", subscriber.GetSubscriberID())
	}

	l.subscribers[subscriber.GetSubscriberID()] = subscriber

	return nil
}

func (l *Listener) GetSubscribers() []ListenerSubscriber {
	l.subscribersLock.RLock()
	defer l.subscribersLock.RUnlock()
	var subscribers []ListenerSubscriber
	for _, subscriber := range l.subscribers {
		subscribers = append(subscribers, subscriber)
	}

	return subscribers
}

func (l *Listener) listenKeyExpiredHandler(msg []byte) {

	var (
		data DataListenKeyExpired
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.ListenKeyExpiredHandler(data)
		}
	}()

}

func (l *Listener) accountUpdateHandler(msg []byte) {

	var (
		data DataAccountUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.AccountBalanceUpdateHandler(data)
			subscriber.AccountPositionUpdateHandler(data)
		}
	}()

}

func (l *Listener) marginCallHandler(msg []byte) {

	var (
		data DataMarginCall
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.MarginCallHandler(data)
		}
	}()

}

func (l *Listener) orderTradeUpdateHandler(msg []byte) {

	var (
		data DataOrderUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.OrderUpdateHandler(data)
		}
	}()

}

func (l *Listener) accountConfigUpdateHandler(msg []byte) {

	var (
		data DataAccountConfigUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.AccountConfigUpdateHandler(data)
		}
	}()

}

func (l *Listener) strategyUpdateHandler(msg []byte) {

	var (
		data DataStrategyUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.StrategyUpdateHandler(data)
		}
	}()

}

func (l *Listener) gridUpdateHandler(msg []byte) {

	var (
		data DataGridUpdate
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.GridUpdateHandler(data)
		}
	}()

}

func (l *Listener) conditionalOrderTriggerRejectHandler(msg []byte) {

	var (
		data DataConditionalOrderTriggerReject
		err  error
	)

	if err = json.Unmarshal(msg, &data); err != nil {
		l.errorHandler(err)
		return
	}

	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.ConditionalOrderTriggerRejectHandler(data)
		}
	}()

}

func (l *Listener) undefinedTypeHandler(msg []byte) {
	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.UndefinedTypeHandler(msg)
		}
	}()
}

func (l *Listener) errorHandler(err error) {
	go func() {
		var subscribers []ListenerSubscriber = l.GetSubscribers()
		for _, subscriber := range subscribers {
			subscriber.ErrorHandler(err)
		}
	}()
}

func (l *Listener) closeHandler() {

	if err := l.Close(); err != nil {
		l.errorHandler(err)
	}

}

func (l *Listener) Close() error {

	if l.wsConn == nil {
		return nil
	}

	if err := l.wsConn.Close(); err != nil {
		return err
	}

	l.isRunning = false

	return nil

}
