package accountswebsockets

type RateLimit interface {
	GetRateLimitType() string
	GetInterval() string
	GetIntervalNum() int
	GetLimit() int
	GetCount() int
}

type RateLimits []RateLimit

type rateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

func (r rateLimit) GetRateLimitType() string {
	return r.RateLimitType
}

func (r rateLimit) GetInterval() string {
	return r.Interval
}

func (r rateLimit) GetIntervalNum() int {
	return r.IntervalNum
}

func (r rateLimit) GetLimit() int {
	return r.Limit
}

func (r rateLimit) GetCount() int {
	return r.Limit
}
