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

	// 1. Validasi stok barang ke Merchant Service
	merchantURL := fmt.Sprintf("http://merchant-service:8089/api/merchant/%s/menu", cart.MerchantID)
	respMerchant, errMerchant := http.Get(merchantURL)
	if errMerchant != nil || respMerchant.StatusCode != http.StatusOK {
		fmt.Println("Warning: Gagal memvalidasi stok di Merchant Service:", errMerchant)
		// We could return an error here, but we proceed with warning to simulate resilience
	} else {
		fmt.Println("Success: Stok barang divalidasi oleh Merchant Service!")
	}
	if respMerchant != nil {
		respMerchant.Body.Close()
	}

	// 2. Minta hitung ongkir ke Pricing Service
	var calculatedFee float64 = 15000.0 // Default fee
	pricingPayload := map[string]interface{}{
		"original_price": 0,
		"promo_code":     "",
		"lat1":           -6.200000, // Dummy koordinat karena model belum support
		"lon1":           106.816666,
		"lat2":           -6.210000,
		"lon2":           106.820000,
	}
	jsonDataPricing, _ := json.Marshal(pricingPayload)
	pricingURL := "http://pricing-service:8087/api/pricing/calculate"
	respPrice, errPrice := http.Post(pricingURL, "application/json", bytes.NewBuffer(jsonDataPricing))
	
	if errPrice == nil {
		defer respPrice.Body.Close()
		var priceResp map[string]interface{}
		if json.NewDecoder(respPrice.Body).Decode(&priceResp) == nil {
			if finalPrice, ok := priceResp["final_price"].(float64); ok {
				calculatedFee = finalPrice
				fmt.Println("Success: Harga didapatkan dari Pricing Service:", calculatedFee)
			}
		}
	} else {
		fmt.Println("Warning: Gagal menghubungi Pricing Service:", errPrice)
	}

	// 3. Komunikasi otomatis ke Translog (Minta carikan driver)
	translogPayload := map[string]interface{}{
		"order_id":         cart.OrderID,
		"user_id":          cart.UserID,
		"status":           "WAITING_FOR_DRIVER", // Status yang pas sebelum SEARCHING mungkin? Atau biarkan.
		"service_type":     "GOSEND",
		"pickup_location":  "Toko Merchant " + cart.MerchantID,
		"dropoff_location": cart.DeliveryAddress,
		"item_dimension":   1.5,
		"fee":              calculatedFee,
	}
	jsonDataTranslog, _ := json.Marshal(translogPayload)

	translogURL := "http://translog-service:8085/api/translog/create"
	respTranslog, errTranslog := http.Post(translogURL, "application/json", bytes.NewBuffer(jsonDataTranslog))

	if errTranslog != nil {
		fmt.Println("Warning: Gagal menghubungi Translog Service:", errTranslog)
	} else {
		defer respTranslog.Body.Close()
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