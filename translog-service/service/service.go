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
    UpdateDeliveryStatus(orderID, status string, fee float64, userID string) error
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
	// 1. Simpan ke Database Translog
	err := s.Repo.SaveOrder(order)
	if err != nil {
		return nil, err
	}

	// 2. OTOMATIS MEMANGGIL PRICING LALU DISPATCH
	if order.Status == "SEARCHING" {
		go func() {
			// A. Nembak ke Pricing Service (Menghitung ongkir)
			pricingPayload := map[string]interface{}{
				"original_price": 0,
				"promo_code":     "",
				"lat1":           -6.200000, // Dummy coordinates as model only stores string address
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
						order.Fee = finalPrice
						// Here we could update the DB with the new Fee
						fmt.Println("[SUCCESS] Harga didapatkan dari Pricing Service:", order.Fee)
					}
				}
			} else {
				fmt.Println("[WARNING] Gagal menghubungi Pricing Service:", errPrice)
			}

			// B. Nembak ke Dispatch Service
			dispatchPayload := map[string]interface{}{
				"order_id":        order.OrderID,
				"service_type":    order.ServiceType,
				"pickup_location": order.PickupLocation,
			}
			jsonDataDispatch, _ := json.Marshal(dispatchPayload)
			
			dispatchURL := "http://dispatch-service-service:8003/api/dispatch/find" 
			respDispatch, errDispatch := http.Post(dispatchURL, "application/json", bytes.NewBuffer(jsonDataDispatch))
			
			if errDispatch != nil {
				fmt.Println("[WARNING] Gagal menghubungi Dispatch Service:", errDispatch)
			} else {
				respDispatch.Body.Close()
				fmt.Println("[SUCCESS] Perintah pencarian driver dikirim ke Dispatch!")
			}
		}()
	}

	return order, nil
}

func (s *TranslogServiceImpl) UpdateDeliveryStatus(orderID, status string, fee float64, userID string) error {
	fmt.Printf("[TRANSLOG] Status resi %s di-update menjadi: %s\n", orderID, status)

	// JIKA DRIVER DITEMUKAN: PANGGIL COMMUNICATION SERVICE (Buka Chat)
	if status == "DRIVER_FOUND" {
		go func() {
			chatPayload := map[string]string{"order_id": orderID, "user_id": userID}
			jsonData, _ := json.Marshal(chatPayload)
			http.Post("http://communication-service:8009/api/chat/init", "application/json", bytes.NewBuffer(jsonData))
			fmt.Println("[SUCCESS] Room Chat Driver-User telah dibuat!")
		}()
	}

	// JIKA SELESAI: PANGGIL FINANCE & SHOP ORDER
	if status == "COMPLETED" || status == "DELIVERED" {
		go func() {
			// A. Potong Saldo (Finance Service)
			financePayload := map[string]interface{}{
				"user_id": userID,
				"amount":  fee,
				"action":  "DEDUCT",
			}
			jsonDataFinance, _ := json.Marshal(financePayload)
			http.Post("http://finance-service:8086/api/wallet/transaction", "application/json", bytes.NewBuffer(jsonDataFinance))
			fmt.Println("[SUCCESS] Instruksi potong saldo dikirim ke Finance!")

			// B. Lapor ke Shop Order (Khusus GoMart/GoSend)
			shopPayload := map[string]string{"order_id": orderID, "status": "COMPLETED"}
			jsonDataShop, _ := json.Marshal(shopPayload)
			http.Post("http://shop-order-service:8084/api/order/update-status", "application/json", bytes.NewBuffer(jsonDataShop))
			fmt.Println("[SUCCESS] Update status pesanan selesai ke Shop Order!")
		}()
	}

	return nil
}