package service

import (
	"encoding/json"
	"finance-service/model"
	"finance-service/repository"
	"fmt"
	"net/http"
)

type WalletService struct {
	repo       repository.WalletRepository
	pricingURL string
}

func NewWalletService(r repository.WalletRepository, pURL string) *WalletService {
	return &WalletService{repo: r, pricingURL: pURL}
}

func (s *WalletService) ProcessTopUp(userID string, amount float64, promoCode string) (float64, error) {
	finalAmount := amount
	
	// Integrasi ke Pricing Service
	if promoCode != "" {
		url := fmt.Sprintf("%s/api/pricing/calculate?amount=%f&promo=%s", s.pricingURL, amount, promoCode)
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			var pr model.PricingResponse
			json.NewDecoder(resp.Body).Decode(&pr)
			finalAmount = pr.FinalAmount
			resp.Body.Close()
		}
	}

	// Update Saldo
	currentBalance, _ := s.repo.GetBalance(userID)
	newBalance := currentBalance + finalAmount
	s.repo.UpdateBalance(userID, newBalance)

	return newBalance, nil
}