package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"translog-service/model"
	"translog-service/repository"
)

type TranslogService interface {
	CreateTransportOrder(order *model.TransportOrder) (*model.TransportOrder, error)
	ValidateStatusTransition(currentStatus, newStatus string) error
	UpdateDeliveryStatus(orderID, status string) error
}

type TranslogServiceImpl struct {
	Repo repository.TranslogRepository
}

func NewTranslogService(repo repository.TranslogRepository) TranslogService {
	return &TranslogServiceImpl{Repo: repo}
}

func (s *TranslogServiceImpl) ValidateStatusTransition(currentStatus, newStatus string) error {
	if currentStatus == "SEARCHING" && newStatus == "COMPLETED" {
		return fmt.Errorf("pesanan tidak bisa langsung COMPLETED dari SEARCHING")
	}
	return nil
}

func (s *TranslogServiceImpl) CreateTransportOrder(order *model.TransportOrder) (*model.TransportOrder, error) {
	err := s.Repo.SaveOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *TranslogServiceImpl) UpdateDeliveryStatus(orderID, status string) error {
	fmt.Printf("[TRANSLOG] Status resi %s di-update menjadi: %s\n", orderID, status)

	if status == "DELIVERED" {
		payload := map[string]string{
			"order_id": orderID,
			"status":   "COMPLETED",
		}
		jsonData, _ := json.Marshal(payload)

		shopURL := "http://shop-order-service:8084/api/order/update-status"
		resp, err := http.Post(shopURL, "application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Warning: Gagal melapor ke Toko:", err)
		} else {
			defer resp.Body.Close()
			fmt.Println("Success: Otomatis melapor ke Toko bahwa pesanan Selesai!")
		}
	}
	return nil
}