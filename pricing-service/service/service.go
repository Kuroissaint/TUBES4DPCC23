package service

import "pricing-service/repository"

type PricingService struct {
	repo repository.PromoRepository
}

func NewPricingService(r repository.PromoRepository) *PricingService {
	return &PricingService{repo: r}
}

func (s *PricingService) CalculateFinalPrice(originalPrice int, promoCode string) (int, error) {
	if promoCode == "" {
		return originalPrice, nil
	}

	discount, err := s.repo.GetDiscountByCode(promoCode)
	if err != nil {
		return 0, err 
	}

	finalPrice := originalPrice - (originalPrice * discount / 100)
	
	return finalPrice, nil
}
