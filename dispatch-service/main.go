package main

import (
	"fmt"
	"net/http"

	"dispatch-service/handler"
	"dispatch-service/service"
)

func main() {
	svc := service.NewDispatchService(nil)
	hdl := handler.NewDispatchHandler(svc)

	http.HandleFunc("/api/dispatch/find", hdl.FindDriverHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})

	fmt.Println("Dispatch Service running on :8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}