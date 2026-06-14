package model

type TransportOrder struct {
	OrderID       string  `json:"order_id"`
	UserID        string  `json:"user_id"`
	Status        string  `json:"status"`
	ServiceType   string  `json:"service_type"`
	ItemDimension float64 `json:"item_dimension,omitempty"`
}
