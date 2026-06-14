package main

import (
	"fmt"
	"net/http"

	"finance-service/handler"
	"finance-service/repository"
	"finance-service/service"
)

type dummyWalletRepo struct{}

func (r *dummyWalletRepo) GetBalance(userID string) (int, error) {
	return 100000, nil // dummy initial balance
}

func (r *dummyWalletRepo) UpdateBalance(userID string, newBalance int) error {
	return nil
}

func main() {
	var repo repository.WalletRepository = &dummyWalletRepo{}
	svc := service.NewWalletService(repo)
	hdl := handler.NewWalletHandler(svc)

	http.HandleFunc("/api/wallet/topup", hdl.TopUpHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Finance service is healthy"))
	})

	fmt.Println("Finance Service running on :8086")
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}