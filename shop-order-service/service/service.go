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
}

type ShopOrderServiceImpl struct {
	Repo repository.ShopOrderRepository
}

func NewShopOrderService(repo repository.ShopOrderRepository) ShopOrderService {
	return &ShopOrderServiceImpl{Repo: repo}
}

func (s *ShopOrderServiceImpl) CreateShoppingOrder() (*model.ShoppingCart, error) {
	// 1. Buat pesanan baru
	cart := &model.ShoppingCart{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MerchantID: uuid.New().String(),
		Items:      []string{"Kopi Gula Aren", "Roti Bakar"},
		Status:     "PAID", // Status diubah jadi PAID biar logis buat dikirim
	}

	// 2. Simpan ke database Supabase (Tabel shop_orders)
	err := s.Repo.SaveCart(cart.OrderID, cart.UserID, cart.MerchantID, cart.Items, cart.Status)
	if err != nil {
		return nil, err
	}

	// 3. --- FITUR KOMUNIKASI OTOMATIS KE TRANSLOG ---
	// Siapkan data pengiriman dengan membawa order_id yang sama
	translogPayload := map[string]interface{}{
		"order_id":       cart.OrderID,
		"user_id":        cart.UserID,
		"status":         "WAITING_FOR_DRIVER",
		"service_type":   "REGULAR",
		"item_dimension": 1.5,
	}
	jsonData, _ := json.Marshal(translogPayload)

	// Tembak API Translog Service di port 8085
	translogURL := "http://localhost:8085/api/translog/create"
	resp, errHTTP := http.Post(translogURL, "application/json", bytes.NewBuffer(jsonData))
	
	if errHTTP != nil {
		fmt.Println("Warning: Gagal menghubungi Translog Service:", errHTTP)
	} else {
		defer resp.Body.Close()
		fmt.Println("Success: Memanggil Translog secara otomatis untuk Order ID:", cart.OrderID)
	}
	// ------------------------------------------------

	return cart, nil
}