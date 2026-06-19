package main

import (
	"database/sql"
	"finance-service/handler"
	"finance-service/repository"
	"finance-service/service"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Wajib di-import agar driver Postgres berjalan
)

func main() {
	// 1. Setup Koneksi Database
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		// Default untuk run di localhost (bukan di dalam cluster)
		dbURL = "postgres://kantin_admin:kantin123@localhost:5432/finance_service_db?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Gagal membuka koneksi DB:", err)
	}
	defer db.Close()

	// Ping untuk memastikan database benar-benar bisa diakses
	if err := db.Ping(); err != nil {
		log.Fatal("Database tidak dapat dijangkau:", err)
	}
	fmt.Println("Berhasil terhubung ke database PostgreSQL (finance_service_db)!")

	// 2. Setup Pricing Service URL
	pricingURL := os.Getenv("PRICING_SERVICE_URL")
	if pricingURL == "" {
		pricingURL = "http://localhost:8081" // Port default pricing-service lokal
	}

	// 3. Inisialisasi Layer (Sekarang pakai NewSqlWalletRepo)
	repo := repository.NewSqlWalletRepo(db)
	svc := service.NewWalletService(repo, pricingURL)
	hdl := handler.NewWalletHandler(svc)

	// 4. Routing
	http.HandleFunc("/api/wallet/topup", hdl.TopUpHandler)

	log.Println("Finance Service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}