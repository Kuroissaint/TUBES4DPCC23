package main

import (
	"fmt"
	"net/http"

	"translog-service/handler"
	"translog-service/repository"
	"translog-service/service"
)

type dummyTranslogRepo struct{}

func (r *dummyTranslogRepo) SaveOrder(orderID, userID, status, serviceType string, itemDimension float64) error {
	return nil
}

func main() {
	var repo repository.TranslogRepository = &dummyTranslogRepo{}
	svc := service.NewTranslogService(repo)
	hdl := handler.NewTranslogHandler(svc)

	http.HandleFunc("/api/translog/create", hdl.CreateTransportOrderHandler)

	fmt.Println("Translog Service running on :8085")
	http.ListenAndServe(":8085", nil)
}
