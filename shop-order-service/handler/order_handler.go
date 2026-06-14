package handler

import (
	"encoding/json"
	"net/http"

	"shop-order-service/service"
)

type OrderHandler struct {
	orderService service.ShopOrderService
}

func NewOrderHandler(os service.ShopOrderService) *OrderHandler {
	return &OrderHandler{orderService: os}
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cart, err := h.orderService.CreateShoppingOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   cart,
	})
}
