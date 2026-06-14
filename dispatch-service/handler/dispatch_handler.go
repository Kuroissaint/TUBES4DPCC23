package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dispatch-service/service"
)

type DispatchHandler struct {
	dispatchService *service.DispatchService
}

func NewDispatchHandler(ds *service.DispatchService) *DispatchHandler {
	return &DispatchHandler{dispatchService: ds}
}

func (h *DispatchHandler) FindDriverHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")

	lat, _ := strconv.ParseFloat(latStr, 64)
	lng, _ := strconv.ParseFloat(lngStr, 64)

	driver, err := h.dispatchService.FindNearestDriver(lat, lng)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"driver": driver,
	})
}
