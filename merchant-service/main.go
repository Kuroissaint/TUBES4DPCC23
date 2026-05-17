package main

import (
	"fmt"
	"net/http"

	"merchant-service/handler"
	"merchant-service/repository"
	"merchant-service/service"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Inisialisasi Repository dengan nil (sesuai pola sementara sebelum connect DB asli)
	repo := repository.NewMerchantRepository(nil)

	// 2. Inisialisasi Service dan inject Repository ke dalamnya
	svc := service.NewMerchantService(repo)

	// 3. Inisialisasi Handler dan inject Service ke dalamnya
	hdl := handler.NewMerchantHandler(svc)

	// 4. Routing HTTP Endpoint
	http.HandleFunc("/api/merchants/register", hdl.RegisterMerchantHandler)

	// 5. Jalankan server HTTP di port 8081 (supaya tidak bentrok dengan account-service)
	fmt.Println("Merchant Service running on :8081")

	http.ListenAndServe(":8081", nil)
}