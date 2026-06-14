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

	http.HandleFunc("/api/location/distance", hdl.CalculateDistanceHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"Location Service is running"}`))
	})

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