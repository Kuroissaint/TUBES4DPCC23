package model

type TopUpRequest struct {
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	PromoCode string  `json:"promo_code,omitempty"`
}

type PricingResponse struct {
	FinalAmount float64 `json:"final_amount"`
}