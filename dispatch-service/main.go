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

	// Daftarkan endpoint operasional penanganan orderan masuk
	http.HandleFunc("/api/dispatch/orders", hdl.CreateOrderHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Dispatch Service is running smoothly"}`))
	})

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8003" // Kita jalankan di port 8003 agar tidak bentrok dengan location-service
	}

	fmt.Printf("Dispatch Service running on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}