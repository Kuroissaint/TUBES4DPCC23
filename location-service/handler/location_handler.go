package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"location-service/service"
)

// Struct untuk membaca body JSON dari driver saat update lokasi
type UpdateLocationRequest struct {
	DriverID  string  `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LocationHandler struct {
	locationService *service.LocationService
	// storageMu dan driverStorage dihapus karena digantikan MongoDB
}

func NewLocationHandler(ls *service.LocationService) *LocationHandler {
	return &LocationHandler{
		locationService: ls,
	}
}

// 1. Endpoint Asli: Menghitung jarak antara dua koordinat (GET)
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid coordinates"})
		return
	}

	distance := h.locationService.CalculateDistance(lat1, lon1, lat2, lon2)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"distance": distance,
	})
}

// 2. Endpoint Baru: Menerima & memperbarui lokasi koordinat driver ke MongoDB (POST)
func (h *LocationHandler) UpdateLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req UpdateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")

	if err != nil || req.DriverID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body atau driver_id kosong"})
		return
	}

	if !h.locationService.ValidateCoordinates(req.Latitude, req.Longitude) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid coordinates"})
		return
	}

	// === SEKARANG SIMPAN KE MONGODB VIA SERVICE ===
	err = h.locationService.SaveDriverLocation(r.Context(), req.DriverID, req.Latitude, req.Longitude)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "gagal menyimpan koordinat ke mongodb: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Lokasi driver " + req.DriverID + " berhasil diperbarui di MongoDB",
	})
}

// 3. Endpoint Baru: Mencari driver terdekat dari koordinat orderan via MongoDB Geospatial (GET)
func (h *LocationHandler) GetNearbyDriversHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	w.Header().Set("Content-Type", "application/json")

	if !h.locationService.ValidateCoordinates(lat, lon) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid coordinates"})
		return
	}

	type NearbyDriverInfo struct {
		DriverID       string  `json:"driver_id"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		DistanceMeters float64 `json:"distance_meters"`
	}

	var nearbyDrivers []NearbyDriverInfo

	// === AMBIL DATA DARI MONGODB VIA SERVICE (Radius 5000 meter) ===
	driversFromDB, err := h.locationService.FindNearbyDrivers(r.Context(), lat, lon, 5000)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "gagal mengambil data dari database"})
		return
	}

	// Mapping hasil query database ke struct Response
	for _, doc := range driversFromDB {
		// Mengingat urutan koordinat GeoJSON MongoDB adalah [lon, lat]
		dLat := doc.Location.Coordinates[1]
		dLon := doc.Location.Coordinates[0]

		// Hitung jarak matematis presisi menggunakan rumus Haversine bawaan service
		distance := h.locationService.CalculateDistance(lat, lon, dLat, dLon)

		nearbyDrivers = append(nearbyDrivers, NearbyDriverInfo{
			DriverID:       doc.DriverID,
			Latitude:       dLat,
			Longitude:      dLon,
			DistanceMeters: distance,
		})
	}

	// Jika MongoDB masih kosong, berikan fallback mock driver
	if len(nearbyDrivers) == 0 {
		mockDistance := h.locationService.CalculateDistance(lat, lon, lat+0.001, lon+0.001)
		nearbyDrivers = append(nearbyDrivers, NearbyDriverInfo{
			DriverID:       "DRV-MOCK-001",
			Latitude:       lat + 0.001,
			Longitude:      lon + 0.001,
			DistanceMeters: mockDistance,
		})
	}

	// Sorting dari yang terdekat berdasarkan DistanceMeters
	sort.Slice(nearbyDrivers, func(i, j int) bool {
		return nearbyDrivers[i].DistanceMeters < nearbyDrivers[j].DistanceMeters
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"drivers": nearbyDrivers,
	})
}