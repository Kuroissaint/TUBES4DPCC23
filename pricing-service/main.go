package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"pricing-service/handler"
	"pricing-service/repository"
	"pricing-service/service"

	_ "github.com/lib/pq" // Driver postgres
)

func main() {
	// 1. Setup Koneksi Database Pricing
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://kantin_admin:kantin123@localhost:5432/pricing_service_db?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Gagal membuka koneksi DB:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database tidak dapat dijangkau:", err)
	}
	fmt.Println("Berhasil terhubung ke database PostgreSQL (pricing_service_db)!")

	// 2. Inisialisasi Layer dengan Repository SQL
	repo := repository.NewSqlPromoRepo(db)
	svc := service.NewPricingService(repo)
	hdl := handler.NewPricingHandler(svc)

	// 3. Routing
	http.HandleFunc("/api/pricing/calculate", hdl.CalculatePriceHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pricing service is healthy"))
	})

	fmt.Println("Pricing and Promo Service running on :8087")
	err = http.ListenAndServe(":8087", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}