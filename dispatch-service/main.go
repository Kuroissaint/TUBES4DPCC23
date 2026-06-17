package main

import (
	"fmt"
	"net/http"
	"os"

	"dispatch-service/handler"
	"dispatch-service/service"
)

func main() {
	svc := service.NewDispatchService(nil)
	hdl := handler.NewDispatchHandler(svc)

	// 1. DAFTARKAN RUTE SPESIFIK DENGAN PREFIX (Sama seperti trik location-service)
	// Karena kodingan asli yang diterima lewat jaringan masih membawa kata /dispatch
	http.HandleFunc("/dispatch/api/dispatch/orders", hdl.CreateOrderHandler)

	// 2. RUTE CATCH-ALL (Hanya merespon jika path pas "/dispatch" atau "/dispatch/")
	http.HandleFunc("/dispatch/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/dispatch" && r.URL.Path != "/dispatch/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Dispatch Service is running smoothly"}`))
	})

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8003" 
	}

	fmt.Printf("Dispatch Service running on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}