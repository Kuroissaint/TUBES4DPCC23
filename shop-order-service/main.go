package main

import (
	"fmt"
	"net/http"

	"shop-order-service/handler"
	"shop-order-service/repository"
	"shop-order-service/service"
)

type dummyOrderRepo struct{}

func (r *dummyOrderRepo) SaveCart(orderID, userID, merchantID string, items []string, status string) error {
	return nil
}

func main() {
	var repo repository.ShopOrderRepository = &dummyOrderRepo{}
	svc := service.NewShopOrderService(repo)
	hdl := handler.NewOrderHandler(svc)

	http.HandleFunc("/api/order/create", hdl.CreateOrderHandler)

	fmt.Println("Shop-Order Service running on :8084")
	http.ListenAndServe(":8084", nil)
}
