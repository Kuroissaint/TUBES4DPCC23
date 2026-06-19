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

	// GUNAKAN MULTIPLEXER (MUX) BAWAAN BIAR TIDAK SALING TABRAKAN
	mux := http.NewServeMux()

	// 1. Rute utama untuk membuat order
	mux.HandleFunc("/dispatch/order", hdl.CreateOrderHandler)

	// 2. Rute pembantu untuk cek status pod (Hanya merespon jika pas "/dispatch" atau "/dispatch/")
	mux.HandleFunc("/dispatch/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/dispatch" && r.URL.Path != "/dispatch/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Dispatch Service Handler Terpasang dengan Rute Baru!"}`))
	})

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8003"
	}

	fmt.Printf("Dispatch Service running on :%s\n", port)
	// Masukkan mux ke dalam ListenAndServe
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}