package accountswebsockets

import (
	"github.com/Lornzo/gobinance/usdmfutures/usdmfutureswebsockettypes"
)

type ErrorSubscriber interface {
	GetID() string
	UpdateError(err error)
}

type errorSubscribers interface {
	UpdateError(err error)
	Subscribe(subscriber usdmfutureswebsockettypes.ErrorSubscriber) error
	UnSubscribe(subscriber usdmfutureswebsockettypes.ErrorSubscriber) error
}
