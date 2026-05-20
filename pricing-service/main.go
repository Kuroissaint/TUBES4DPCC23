package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Pricing and Promo Service Running...")

	// Handler sederhana untuk cek status health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pricing service is healthy"))
	})

	// Menahan aplikasi agar tetap hidup di port 8081 (atau sesuaikan port k8s kamu)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}