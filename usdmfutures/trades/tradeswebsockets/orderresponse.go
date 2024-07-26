package tradeswebsockets

type orderResponse struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Result Order  `json:"result"`
}
