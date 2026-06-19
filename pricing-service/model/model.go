package model

type CalculatePriceRequest struct {
	OriginalPrice int    `json:"original_price"`
	PromoCode     string `json:"promo_code"`
}
