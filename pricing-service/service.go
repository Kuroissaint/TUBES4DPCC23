package main

type PricingService struct {
	repo PromoRepository
}

func NewPricingService(r PromoRepository) *PricingService {
	return &PricingService{repo: r}
}

// CalculateFinalPrice adalah fungsi yang nanti akan kamu lengkapi kodenya
func (s *PricingService) CalculateFinalPrice(originalPrice int, promoCode string) (int, error) {
	// TODO: Nanti coding logika aslinya di sini
	return 0, nil
}