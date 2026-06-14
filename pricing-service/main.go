package main

import (
	"fmt"
	"net/http"

	"pricing-service/handler"
	"pricing-service/repository"
	"pricing-service/service"
)

type dummyPricingRepo struct{}

func (r *dummyPricingRepo) GetDiscountByCode(promoCode string) (int, error) {
	return 10, nil // 10% discount
}

func main() {
	var repo repository.PromoRepository = &dummyPricingRepo{}
	svc := service.NewPricingService(repo)
	hdl := handler.NewPricingHandler(svc)

	http.HandleFunc("/api/pricing/calculate", hdl.CalculatePriceHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pricing service is healthy"))
	})

	fmt.Println("Pricing and Promo Service running on :8087")
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}