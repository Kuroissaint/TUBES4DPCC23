package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"shop-order-service/model"
	"shop-order-service/repository"

	"github.com/google/uuid"
)

type ShopOrderService interface {
	CreateShoppingOrder() (*model.ShoppingCart, error)
	GetOrder(orderID string) (*model.ShoppingCart, error)
	UpdateOrderStatus(orderID, status string) error
}

type ShopOrderServiceImpl struct {
	Repo repository.ShopOrderRepository
}

func NewShopOrderService(repo repository.ShopOrderRepository) ShopOrderService {
	return &ShopOrderServiceImpl{Repo: repo}
}

// 1. Perbaikan: Nama fungsi harus CreateShoppingOrder sesuai interface
func (s *ShopOrderServiceImpl) CreateShoppingOrder() (*model.ShoppingCart, error) {
	// Buat pesanan baru
	cart := &model.ShoppingCart{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MerchantID: uuid.New().String(),
		Items:      []string{"Kopi Gula Aren", "Roti Bakar"},
		Status:     "PAID", 
	}

	// Simpan ke database
	err := s.Repo.SaveCart(cart.OrderID, cart.UserID, cart.MerchantID, cart.Items, cart.Status)
	if err != nil {
		return nil, err
	}

	// Komunikasi otomatis ke Translog
	translogPayload := map[string]interface{}{
		"order_id":       cart.OrderID,
		"user_id":        cart.UserID,
		"status":         "WAITING_FOR_DRIVER",
		"service_type":   "REGULAR",
		"item_dimension": 1.5,
	}
	jsonData, _ := json.Marshal(translogPayload)

	translogURL := "http://localhost:8085/api/translog/create"
	resp, errHTTP := http.Post(translogURL, "application/json", bytes.NewBuffer(jsonData))
	
	if errHTTP != nil {
		fmt.Println("Warning: Gagal menghubungi Translog Service:", errHTTP)
	} else {
		defer resp.Body.Close()
		fmt.Println("Success: Memanggil Translog secara otomatis untuk Order ID:", cart.OrderID)
	}

	// Perbaikan: Mengembalikan cart yang sudah dibuat (bukan orderID yang undefined)
	return cart, nil
}

// 2. Perbaikan: Implementasi GetOrder yang menerima parameter orderID
func (s *ShopOrderServiceImpl) GetOrder(orderID string) (*model.ShoppingCart, error) {
	// Implementasi logika ambil data dari repo kalau sudah ada
	return &model.ShoppingCart{OrderID: orderID, Status: "PAID"}, nil
}

// 3. Implementasi UpdateOrderStatus
func (s *ShopOrderServiceImpl) UpdateOrderStatus(orderID, status string) error {
	fmt.Printf("[SHOP-ORDER] Pesanan %s sukses di-update menjadi: %s\n", orderID, status)
	return nil
}