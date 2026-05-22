package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "shop-order-service/shoporder" // Import package baru
)

func createShoppingHandler(svc shoporder.ShopOrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cart, err := svc.CreateShoppingOrder()
		if err != nil {
			http.Error(w, "Gagal memproses pesanan", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(cart)
	}
}

func main() {
	repo := shoporder.NewShopOrderRepository()
	svc := shoporder.NewShopOrderService(repo)

	http.HandleFunc("/api/orders/shopping", createShoppingHandler(svc))
	log.Fatal(http.ListenAndServe(":8009", nil))
}