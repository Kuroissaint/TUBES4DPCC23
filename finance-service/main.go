package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Finance Service Running...")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Finance service is healthy"))
	})

	// Menahan aplikasi agar tetap hidup di port 8082
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}