package main

type PricingService struct {
	repo PromoRepository
}

func NewPricingService(r PromoRepository) *PricingService {
	return &PricingService{repo: r}
}

func (s *PricingService) CalculateFinalPrice(originalPrice int, promoCode string) (int, error) {
	// Jika tidak ada promo, kembalikan harga asli
	if promoCode == "" {
		return originalPrice, nil
	}

	// Ambil persentase diskon dari repository (bisa via mock atau DB asli)
	discount, err := s.repo.GetDiscountByCode(promoCode)
	if err != nil {
		return 0, err // Jika error (misal promo tidak valid), return error
	}

	// Hitung harga akhir
	finalPrice := originalPrice - (originalPrice * discount / 100)
	
	return finalPrice, nil
}