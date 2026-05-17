package main

import (
	"account-service/handler"
	"account-service/repository"
	"account-service/service"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main () {
	repo := repository.NewUserRepository(nil)

	svc := service.NewAuthService(repo)

	hdl := handler.NewAuthHandler(svc)

	// Routing
	http.HandleFunc("/api/auth/login", hdl.LoginHandler)

	fmt.Println("Account Service running on :8080")

	http.ListenAndServe(":8080", nil)
}