package main

import (
	"account-service/handler"
	"account-service/repository"
	"account-service/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	//koneksi database supabase//
	connStr := "host=aws-1-ap-northeast-1.pooler.supabase.com port=6543 user=postgres.xjsjnbetglgcaimpckaj password=firdatubes1 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}
	defer db.Close()

	// Cek koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak bisa di-ping:", err)
	}
	fmt.Println("Database terhubung!")

	repo := repository.NewUserRepository(db)
	svc := service.NewAuthService(repo)
	hdl := handler.NewAuthHandler(svc)

	// Routing
	http.HandleFunc("/api/auth/login", hdl.LoginHandler)
	http.HandleFunc("/api/auth/register", hdl.RegisterHandler) //
	// //(untuk driver)
	http.HandleFunc("/api/driver/register", hdl.DriverRegisterHandler)
	//(untuk customer)
	http.HandleFunc("/api/customer/register", hdl.CustomerRegisterHandler)

	fmt.Println("Account Service running on :8081")
	http.ListenAndServe(":8081", nil)
}
