package trades

type PlacedOrder struct {
	ClientOrderID string `json:"clientOrderId,omitempty"`
	OrderID       int64  `json:"orderId,omitempty"`
}
