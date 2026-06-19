package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pricing-service/repository"
)

type PricingService struct {
	repo repository.PromoRepository
}

func NewPricingService(r repository.PromoRepository) *PricingService {
	return &PricingService{repo: r}
}

// CalculateFinalPrice calculates the base price using distance from location-service and applies promo.
func (s *PricingService) CalculateFinalPrice(originalPrice int, promoCode string, lat1, lon1, lat2, lon2 float64) (int, error) {
	// Call location service to calculate distance
	locationURL := fmt.Sprintf("http://location-service-service:8002/location/distance?lat1=%f&lon1=%f&lat2=%f&lon2=%f", lat1, lon1, lat2, lon2)
	resp, err := http.Get(locationURL)
	var distance float64 = 5.0 // fallback distance
	if err == nil {
		defer resp.Body.Close()
		var locResp struct {
			Distance float64 `json:"distance"`
		}
		if json.NewDecoder(resp.Body).Decode(&locResp) == nil {
			distance = locResp.Distance / 1000 // Convert meters to KM
		}
	} else {
		fmt.Println("[WARNING] Gagal menghubungi Location Service:", err)
	}

	// Base fare Rp 2000 per KM
	basePrice := int(distance * 2000)
	if basePrice < 10000 {
		basePrice = 10000 // Minimum fare
	}

	finalPrice := originalPrice + basePrice

	if promoCode != "" {
		discount, err := s.repo.GetDiscountByCode(promoCode)
		if err == nil {
			finalPrice = finalPrice - (finalPrice * discount / 100)
		}
	}
	
	return finalPrice, nil
}
