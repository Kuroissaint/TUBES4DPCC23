package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type StandardResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ShoppingCart struct {
	OrderID    string   `json:"order_id"`
	UserID     string   `json:"user_id"`
	MerchantID string   `json:"merchant_id"`
	Items      []string `json:"items"`
	Status     string   `json:"status"`
}

// Logika Bisnis: Menambah item ke keranjang
func (c *ShoppingCart) AddToCart(item string) {
	c.Items = append(c.Items, item)
}

// Handler HTTP
func createShoppingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	cart := ShoppingCart{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MerchantID: uuid.New().String(),
		Items:      []string{"Kopi Gula Aren", "Roti Bakar"},
		Status:     "AT_STORE",
	}

	response := StandardResponse{
		Status:  "success",
		Data:    cart,
		Message: "Pesanan belanja berhasil dikonfirmasi",
	}
	
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/orders/shopping", createShoppingHandler)
	fmt.Println("Shopping Order Service berjalan di port 8009...")
	log.Fatal(http.ListenAndServe(":8009", nil))
}