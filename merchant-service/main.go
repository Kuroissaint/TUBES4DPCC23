package main

import (
	"database/sql"
	"fmt" //tampilan/print pesan ke terminal
	"log" //print eror dan menghentikan program jika terjadi masalah
	"merchant-service/handler"
	"merchant-service/repository"
	"merchant-service/service"
	"net/http" //agar program dapat menerima request dari luar (agar go bisa menjadi web server)

	_ "github.com/lib/pq"
)

func main() {

	//koneksi database supabase//
	// connStr := "host=aws-1-ap-northeast-1.pooler.supabase.com port=6543 user=postgres.xjsjnbetglgcaimpckaj password=firdatubes1 dbname=postgres sslmode=disable"
	connStr := "postgresql://postgres.xjsjnbetglgcaimpckaj:firdatubes1@aws-1-ap-northeast-1.pooler.supabase.com:6543/postgres?pgbouncer=true"
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

	// 1. Inisialisasi Repository dengan nil (sesuai pola sementara sebelum connect DB asli)
	repo := repository.NewMerchantRepository(db)
	// 2. Inisialisasi Service dan inject Repository ke dalamnya
	svc := service.NewMerchantService(repo)
	// 3. Inisialisasi Handler dan inject Service ke dalamnya
	hdl := handler.NewMerchantHandler(svc)

	// 4. Routing HTTP Endpoint
	http.HandleFunc("/api/merchant/register", hdl.RegisterMerchantHandler)
	http.HandleFunc("/api/merchant/", hdl.MerchantRouterHandler)
	// 5. Jalankan server HTTP di port 8089 (supaya tidak bentrok dengan account-service)
	fmt.Println("Merchant Service running on :8089")
	http.ListenAndServe(":8089", nil)
}
