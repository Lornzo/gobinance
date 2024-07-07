package accountsrests

type RateLimitOrderResponseItem struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
}

type RateLimitOrderResponse []RateLimitOrderResponseItem
