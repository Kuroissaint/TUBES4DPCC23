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

	// Menerima rute utuh dari gerbang depan secara sempurna!
	http.HandleFunc("/dispatch/api/dispatch/orders", hdl.CreateOrderHandler)

	// Rute pembantu penanda pod v99 aktif
	http.HandleFunc("/dispatch/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/dispatch" && r.URL.Path != "/dispatch/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Dispatch Service Handler Terpasang!"}`))
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