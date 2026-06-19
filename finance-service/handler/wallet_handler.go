package handler

import (
	"encoding/json"
	"finance-service/model"
	"finance-service/service"
	"net/http"
)

type WalletHandler struct {
	ws *service.WalletService
}

func NewWalletHandler(ws *service.WalletService) *WalletHandler {
	return &WalletHandler{ws: ws}
}

func (h *WalletHandler) TopUpHandler(w http.ResponseWriter, r *http.Request) {
	var req model.TopUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newBalance, err := h.ws.ProcessTopUp(req.UserID, req.Amount, req.PromoCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"new_balance": newBalance,
	})
}