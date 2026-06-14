package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"location-service/service"
)

type LocationHandler struct {
	locationService *service.LocationService
}

func NewLocationHandler(ls *service.LocationService) *LocationHandler {
	return &LocationHandler{locationService: ls}
}

func (h *LocationHandler) CalculateDistanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	lat1, _ := strconv.ParseFloat(r.URL.Query().Get("lat1"), 64)
	lon1, _ := strconv.ParseFloat(r.URL.Query().Get("lon1"), 64)
	lat2, _ := strconv.ParseFloat(r.URL.Query().Get("lat2"), 64)
	lon2, _ := strconv.ParseFloat(r.URL.Query().Get("lon2"), 64)

	if !h.locationService.ValidateCoordinates(lat1, lon1) || !h.locationService.ValidateCoordinates(lat2, lon2) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid coordinates"})
		return
	}

	distance := h.locationService.CalculateDistance(lat1, lon1, lat2, lon2)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"distance": distance,
	})
}
