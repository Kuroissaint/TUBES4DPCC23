package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"dispatch-service/service"
)

// Request body saat pelanggan membuat orderan baru
type CreateOrderRequest struct {
	OrderID   string  `json:"order_id"`
	PickupLat float64 `json:"pickup_latitude"`
	PickupLon float64 `json:"pickup_longitude"`
}

// Struktur penampung response dari Location Service
type LocationServiceResponse struct {
	Status  string `json:"status"`
	Drivers []struct {
		DriverID       string      `json:"driver_id"`
		DistanceMeters interface{} `json:"distance_meters"`
	} `json:"drivers"`
}

type DispatchHandler struct {
	dispatchService *service.DispatchService
}

func NewDispatchHandler(ds *service.DispatchService) *DispatchHandler {
	return &DispatchHandler{dispatchService: ds}
}

// Endpoint Baru: Menerima orderan dan mencocokkannya dengan driver terdekat (POST)
func (h *DispatchHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")

	if err != nil || req.OrderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body atau order_id kosong"})
		return
	}

	// AMBIL URL LOCATION SERVICE DARI ENVIRONMENT
	locationServiceURL := os.Getenv("LOCATION_SERVICE_URL")
	if locationServiceURL == "" {
		locationServiceURL = "http://localhost:8002" // Default ke port lokal kamu
	}

	// === SINKRONISASI TEMBAKAN ANTAR-SERVICE: Ubah ke path baru /location/nearby ===
	apiURL := fmt.Sprintf("%s/location/nearby?lat=%f&lon=%f", locationServiceURL, req.PickupLat, req.PickupLon)
	resp, err := http.Get(apiURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "gagal menghubungi location-service: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	var locResp LocationServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&locResp); err != nil || len(locResp.Drivers) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"order_id": req.OrderID,
			"status":   "failed",
			"message":  "tidak ada driver terdekat yang tersedia di sekitar lokasi penjemputan",
		})
		return
	}

	// Ambil driver urutan pertama (paling dekat hasilnya karena sudah di-sorting oleh location-service)
	driverTerpilih := locResp.Drivers[0]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":          "success",
		"order_id":        req.OrderID,
		"assigned_driver": driverTerpilih.DriverID,
		"distance_meters": driverTerpilih.DistanceMeters,
		"message":         fmt.Sprintf("Order %s berhasil ditugaskan ke %s", req.OrderID, driverTerpilih.DriverID),
	})
}