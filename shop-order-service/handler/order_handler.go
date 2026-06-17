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

// 1. Handler untuk Create Order (Sudah ada sebelumnya)
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

// 2. Handler untuk Get Order (Ini yang tadi bikin error undefined)
func (h *OrderHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing order ID"})
		return
	}

	cart, _ := h.orderService.GetOrder(orderID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   cart,
	})
}

// Struct bantuan untuk membaca JSON payload update status
type UpdateStatusPayload struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

// 3. Handler untuk Update Status (Menerima laporan dari kurir)
func (h *OrderHandler) UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var payload UpdateStatusPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Panggil service untuk mengupdate status
	err := h.orderService.UpdateOrderStatus(payload.OrderID, payload.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Status pesanan toko berhasil diupdate",
	})
}