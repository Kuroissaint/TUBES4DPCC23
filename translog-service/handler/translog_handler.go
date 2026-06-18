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

type CreateTranslogPayload struct {
	OrderID       string  `json:"order_id"`
	UserID        string  `json:"user_id"`
	Status        string  `json:"status"`
	ServiceType   string  `json:"service_type"`
	ItemDimension float64 `json:"item_dimension"`
}

func (h *TranslogHandler) CreateTransportOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 1. Tangkap JSON dari Body Postman
	var payload CreateTranslogPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON tidak valid"})
		return
	}

	// 2. Eksekusi ke Service
	// CATATAN KRITIS: Sama seperti shop-order, ubah service.go Anda
	// agar CreateTransportOrder menerima parameter dari payload ini.
	order, err := h.translogService.CreateTransportOrder() 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   order, // order harus mengembalikan data yang baru saja disimpan
	})
}

// Tambahan handler agar bukan nano-service
func (h *TranslogHandler) CreateTransportOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var payload CreateTranslogPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON tidak valid"})
		return
	}

	// 1. Petakan DTO ke Model
	orderReq := &model.TransportOrder{
		OrderID:       payload.OrderID,
		UserID:        payload.UserID,
		Status:        payload.Status,
		ServiceType:   payload.ServiceType,
		ItemDimension: payload.ItemDimension,
	}

	// 2. Teruskan ke Service
	order, err := h.translogService.CreateTransportOrder(orderReq) 
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

