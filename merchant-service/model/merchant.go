package model

type MerchantRequest struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type MerchantResponse struct {
	MerchantID int    `json:"merchant_id"`
	Status     string `json:"status"`
}

type Merchant struct {
	ID   int
	Name string
	City string
}