package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Standar Response JSON
type StandardResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type TransportOrder struct {
	OrderID       string  `json:"order_id"`
	UserID        string  `json:"user_id"`
	Status        string  `json:"status"`
	ServiceType   string  `json:"service_type"`
	ItemDimension float64 `json:"item_dimension,omitempty"` // Khusus delivery
}

// Logika Bisnis: Validasi Transisi Status
func ValidateStatusTransition(currentStatus, newStatus string) error {
	if currentStatus == "SEARCHING" && newStatus == "COMPLETED" {
		return fmt.Errorf("pesanan tidak bisa langsung COMPLETED dari SEARCHING")
	}
	return nil
}

// Handler HTTP
func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	newOrder := TransportOrder{
		OrderID:     uuid.New().String(),
		UserID:      uuid.New().String(),
		Status:      "SEARCHING",
		ServiceType: "ride",
	}

	response := StandardResponse{
		Status:  "success",
		Data:    newOrder,
		Message: "Order transportasi berhasil dibuat",
	}
	
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/orders/transport", createOrderHandler)
	fmt.Println("Transport Order Service berjalan di port 8005...")
	log.Fatal(http.ListenAndServe(":8005", nil))
}