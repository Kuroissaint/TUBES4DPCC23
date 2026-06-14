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
