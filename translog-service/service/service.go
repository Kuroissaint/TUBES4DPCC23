package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"translog-service/model"
	"translog-service/repository"
)

type TranslogService interface {
	CreateTransportOrder() (*model.TransportOrder, error)
	ValidateStatusTransition(currentStatus, newStatus string) error
	UpdateDeliveryStatus(orderID, status string) error // <-- Fitur Laporan Kurir
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

func (s *TranslogServiceImpl) CreateTransportOrder() (*model.TransportOrder, error) {
	newOrder := &model.TransportOrder{
		OrderID:       uuid.New().String(),
		UserID:        uuid.New().String(),
		Status:        "SEARCHING",
		ServiceType:   "ride",
		ItemDimension: 0.0,
	}

	err := s.Repo.SaveOrder(newOrder.OrderID, newOrder.UserID, newOrder.Status, newOrder.ServiceType, newOrder.ItemDimension)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

// --- FITUR BARU: KOMUNIKASI BALIK KE SHOP-ORDER ---
func (s *TranslogServiceImpl) UpdateDeliveryStatus(orderID, status string) error {
	// Idealnya di sini memanggil s.Repo.UpdateStatus(...) untuk update ke DB Translog
	fmt.Printf("[TRANSLOG] Status resi %s di-update menjadi: %s\n", orderID, status)

	// Jika barang sudah sampai (DELIVERED), lapor ke toko secara otomatis!
	if status == "DELIVERED" {
		payload := map[string]string{
			"order_id": orderID,
			"status":   "COMPLETED",
		}
		jsonData, _ := json.Marshal(payload)

		// Tembak API Shop Order di port 8084
		shopURL := "http://localhost:8084/api/order/update-status"
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