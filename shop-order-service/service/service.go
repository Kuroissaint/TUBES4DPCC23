package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"shop-order-service/model"
	"shop-order-service/repository"
)

type ShopOrderService interface {
	// 1. Tambahkan parameter cart di interface
	CreateShoppingOrder(cart *model.ShoppingCart) (*model.ShoppingCart, error)
	GetOrder(orderID string) (*model.ShoppingCart, error)
	UpdateOrderStatus(orderID, status string) error
}

type ShopOrderServiceImpl struct {
	Repo repository.ShopOrderRepository
}

func NewShopOrderService(repo repository.ShopOrderRepository) ShopOrderService {
	return &ShopOrderServiceImpl{Repo: repo}
}

// 2. Terima parameter dari Handler, HAPUS hardcode UUID!
func (s *ShopOrderServiceImpl) CreateShoppingOrder(cart *model.ShoppingCart) (*model.ShoppingCart, error) {
	// Simpan ke database menggunakan data dinamis dari Postman
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

	// PERBAIKAN FATAL: Gunakan nama service Kubernetes, bukan localhost!
	translogURL := "http://translog-service:8085/api/translog/create"
	resp, errHTTP := http.Post(translogURL, "application/json", bytes.NewBuffer(jsonData))
	
	if errHTTP != nil {
		fmt.Println("Warning: Gagal menghubungi Translog Service:", errHTTP)
	} else {
		defer resp.Body.Close()
		fmt.Println("Success: Memanggil Translog secara otomatis untuk Order ID:", cart.OrderID)
	}

	return cart, nil
}

// (Biarkan implementasi GetOrder dan UpdateOrderStatus Anda yang lama di bawah sini)
func (s *ShopOrderServiceImpl) GetOrder(orderID string) (*model.ShoppingCart, error) {
	return &model.ShoppingCart{OrderID: orderID, Status: "PAID"}, nil
}

func (s *ShopOrderServiceImpl) UpdateOrderStatus(orderID, status string) error {
	fmt.Printf("[SHOP-ORDER] Pesanan %s sukses di-update menjadi: %s\n", orderID, status)
	return nil
}