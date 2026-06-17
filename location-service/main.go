package main

import (
	"fmt"
	"net/http"
	"os"

	"location-service/handler"
	"location-service/service"
)

func main() {
	svc := service.NewLocationService(nil)
	hdl := handler.NewLocationHandler(svc)

	// 1. DAFTARKAN RUTE SPESIFIK DULU (Wajib di awal agar tidak dibajak rute default)
	http.HandleFunc("/location/api/location/distance", hdl.CalculateDistanceHandler)
	http.HandleFunc("/location/api/location/update", hdl.UpdateLocationHandler)
	http.HandleFunc("/location/api/location/nearby", hdl.GetNearbyDriversHandler)

	// 2. RUTE CATCH-ALL (Hanya merespon jika path pas "/location" atau "/location/")
	http.HandleFunc("/location/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/location" && r.URL.Path != "/location/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Location Service is running smoothly"}`))
	})

	// 3. JALANKAN SERVER
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8002"
	}

	fmt.Printf("Location Service running on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}