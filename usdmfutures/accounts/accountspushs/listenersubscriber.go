package accountspushs

type ListenerSubscriber interface {
	GetSubscriberID() string
	ListenKeyExpiredHandler(data ListenKeyExpired)
	AccountBalanceUpdateHandler(data AccountBalance)
	AccountPositionUpdateHandler(data AccountPosition)
	MarginCallHandler(data MarginCall)
	OrderUpdateHandler(data OrderUpdate)
	AccountConfigUpdateHandler(data AccountConfigUpdate)
	StrategyUpdateHandler(data StrategyUpdate)
	GridUpdateHandler(data GridUpdate)
	ConditionalOrderTriggerRejectHandler(data ConditionalOrderTriggerReject)
	UndefinedTypeHandler(typeData []byte)
	ErrorHandler(err error)
}
