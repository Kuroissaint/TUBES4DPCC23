package handler

import (
	"encoding/json"
	"net/http"

	"translog-service/service"
)

type TranslogHandler struct {
	translogService service.TranslogService
}

func NewTranslogHandler(ts service.TranslogService) *TranslogHandler {
	return &TranslogHandler{translogService: ts}
}

func (h *TranslogHandler) CreateTransportOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	order, err := h.translogService.CreateTransportOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   order,
	})
}

// Tambahan handler agar bukan nano-service
func (h *TranslogHandler) GetTransportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Query().Get("order_id")
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Detail log pengiriman untuk order: " + orderID,
	})
}

type UpdateTranslogPayload struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

// Tambahkan di bagian bawah file:
func (h *TranslogHandler) UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var payload UpdateTranslogPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.translogService.UpdateDeliveryStatus(payload.OrderID, payload.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"message": "Status pengiriman berhasil diupdate",
	})
}