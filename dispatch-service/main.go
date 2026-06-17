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

	// JURUS RUTE GANDA: Gateway motong atau nggak, request PASTI MASUK!
	http.HandleFunc("/api/dispatch/orders", hdl.CreateOrderHandler)
	http.HandleFunc("/dispatch/api/dispatch/orders", hdl.CreateOrderHandler)

	// RUTE HEALTH CHECK / CATCH-ALL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Dispatch Service v99 is RUNNING!"}`))
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