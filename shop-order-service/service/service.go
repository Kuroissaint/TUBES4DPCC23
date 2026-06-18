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

func (s *ShopOrderServiceImpl) CreateShoppingOrder(cart *model.ShoppingCart) (*model.ShoppingCart, error) {
	err := s.Repo.SaveCart(cart)
	if err != nil {
		return nil, err
	}

	// Komunikasi otomatis ke Translog (Menggunakan data dinamis yang baru)
	translogPayload := map[string]interface{}{
		"order_id":         cart.OrderID,
		"user_id":          cart.UserID,
		"status":           "WAITING_FOR_DRIVER",
		"service_type":     "GOSEND",
		"pickup_location":  "Toko Merchant " + cart.MerchantID, // Simulasi alamat toko
		"dropoff_location": cart.DeliveryAddress,               // Alamat dari pemesan
		"item_dimension":   1.5,
		"fee":              15000.00, // Simulasi ongkir tetap untuk GoMart
	}
	jsonData, _ := json.Marshal(translogPayload)

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

func (s *ShopOrderServiceImpl) GetOrder(orderID string) (*model.ShoppingCart, error) {
	return &model.ShoppingCart{OrderID: orderID, Status: "PAID"}, nil
}

func (s *ShopOrderServiceImpl) UpdateOrderStatus(orderID, status string) error {
	fmt.Printf("[SHOP-ORDER] Pesanan %s sukses di-update menjadi: %s\n", orderID, status)
	return nil
}