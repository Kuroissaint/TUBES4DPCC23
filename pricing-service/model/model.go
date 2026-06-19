package model

type CalculatePriceRequest struct {
	OriginalPrice int     `json:"original_price"`
	PromoCode     string  `json:"promo_code"`
	Lat1          float64 `json:"lat1"`
	Lon1          float64 `json:"lon1"`
	Lat2          float64 `json:"lat2"`
	Lon2          float64 `json:"lon2"`
}
