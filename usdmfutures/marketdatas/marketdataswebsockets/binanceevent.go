package marketdataswebsockets

type binanceEvent struct {
	EventName        string `json:"e"`
	EventMilliSecond int64  `json:"E"`
}
