package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Import driver PostgreSQL

	"translog-service/handler"
	"translog-service/repository"
	"translog-service/service"
)

func main() {
	// 1. Ambil URL database dari terminal/sistem
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("ERROR: DATABASE_URL is not set in environment")
	}

	// 2. Buka koneksi ke Supabase PostgreSQL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}
	defer db.Close()

	// 3. Tes koneksi (Ping)
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database. Check your URL/Password:", err)
	}
	fmt.Println("Database connection to Supabase established successfully!")

	// 4. Masukkan DB ke repository
	var repo repository.TranslogRepository = repository.NewTranslogRepository(db)
	svc := service.NewTranslogService(repo)
	hdl := handler.NewTranslogHandler(svc)

	http.HandleFunc("/api/translog/create", hdl.CreateTransportOrderHandler)

	fmt.Println("Translog Service running on :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}