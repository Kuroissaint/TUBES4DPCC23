package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Import driver PostgreSQL

	"shop-order-service/handler"
	"shop-order-service/repository"
	"shop-order-service/service"
)

func main() {
	// 1. Ambil URL database dari terminal/sistem
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("ERROR: DATABASE_URL is not set in environment")
	}

	// 2. Buka koneksi ke Supabase
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database. Check your URL/Password:", err)
	}
	fmt.Println("Database connection to Supabase established successfully!")

	// 3. Setup dependencies
	var repo repository.ShopOrderRepository = repository.NewShopOrderRepository(db)
	svc := service.NewShopOrderService(repo)
	hdl := handler.NewOrderHandler(svc)

	http.HandleFunc("/api/order/create", hdl.CreateOrderHandler)

	fmt.Println("Shop-Order Service running on :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}