package model

type TransportOrder struct {
	OrderID         string   `json:"order_id"`
	UserID          string   `json:"user_id"`
	DriverID        *string  `json:"driver_id,omitempty"` // Nullable
	Status          string   `json:"status"`
	ServiceType     string   `json:"service_type"`
	PickupLocation  string   `json:"pickup_location"`
	DropoffLocation string   `json:"dropoff_location"`
	ItemDimension   *float64 `json:"item_dimension,omitempty"` // Nullable
	Fee             float64  `json:"fee"`
}