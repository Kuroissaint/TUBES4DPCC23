package handler

import (
	"encoding/json"
	"net/http"

	"pricing-service/model"
	"pricing-service/service"
)

type PricingHandler struct {
	pricingService *service.PricingService
}

func NewPricingHandler(ps *service.PricingService) *PricingHandler {
	return &PricingHandler{pricingService: ps}
}

func (h *PricingHandler) CalculatePriceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req model.CalculatePriceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	finalPrice, err := h.pricingService.CalculateFinalPrice(req.OriginalPrice, req.PromoCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      "success",
		"final_price": finalPrice,
	})
}
