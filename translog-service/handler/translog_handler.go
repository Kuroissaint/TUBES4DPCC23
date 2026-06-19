package handler

import (
	"encoding/json"
	"net/http"
	"translog-service/model"
	"translog-service/service"
)

type TranslogHandler struct {
	translogService service.TranslogService
}

func NewTranslogHandler(ts service.TranslogService) *TranslogHandler {
	return &TranslogHandler{translogService: ts}
}

type CreateTranslogPayload struct {
	OrderID         string   `json:"order_id"`
	UserID          string   `json:"user_id"`
	Status          string   `json:"status"`
	ServiceType     string   `json:"service_type"`
	PickupLocation  string   `json:"pickup_location"`
	DropoffLocation string   `json:"dropoff_location"`
	ItemDimension   *float64 `json:"item_dimension"`
	Fee             float64  `json:"fee"`
}

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

	orderReq := &model.TransportOrder{
		OrderID:         payload.OrderID,
		UserID:          payload.UserID,
		Status:          payload.Status,
		ServiceType:     payload.ServiceType,
		PickupLocation:  payload.PickupLocation,
		DropoffLocation: payload.DropoffLocation,
		ItemDimension:   payload.ItemDimension,
		Fee:             payload.Fee,
	}

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

// PERBAIKAN: Tambahkan Fee dan UserID agar bisa diambil dari body request
type UpdateTranslogPayload struct {
	OrderID string  `json:"order_id"`
	Status  string  `json:"status"`
	Fee     float64 `json:"fee"`
	UserID  string  `json:"user_id"`
}

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

	// PERBAIKAN: Mengirim 4 parameter sesuai kebutuhan interface
	err := h.translogService.UpdateDeliveryStatus(payload.OrderID, payload.Status, payload.Fee, payload.UserID)
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