package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"merchant-service/model"
)

type MockMerchantService struct{}

func (m MockMerchantService) RegisterMerchant(req model.MerchantRegisterRequest) (*model.MerchantRegisterResponse, error) {
	return &model.MerchantRegisterResponse{
		UserID:  1,
		Message: "Registrasi merchant berhasil!",
	}, nil
}

func (m MockMerchantService) GetMerchant(id int) (*model.Merchant, error) {
	return &model.Merchant{
		ID:       id,
		NamaToko: "Warung Test",
		Kota:     "Bandung",
	}, nil
}

func (m MockMerchantService) AddMenuItem(merchantID int, req model.MenuItemRequest) (*model.MenuItemResponse, error) {
	return &model.MenuItemResponse{
		ID:      1,
		Message: "Menu berhasil ditambahkan!",
	}, nil
}

func (m MockMerchantService) GetMenu(merchantID int) ([]model.MenuItem, error) {
	return []model.MenuItem{
		{ID: 1, MerchantID: merchantID, Nama: "Nasi Goreng", Harga: 15000},
	}, nil
}

func TestRegisterMerchantHandler(t *testing.T) {
	jsonBody := []byte(`{
		"nama": "Pak Budi",
		"email": "pakbudi@gmail.com",
		"no_hp": "081234567890",
		"password": "password123",
		"nama_toko": "Warung Pak Budi",
		"kota": "Bandung"
	}`)

	req := httptest.NewRequest("POST", "/api/merchant/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	mockService := MockMerchantService{}
	handler := NewMerchantHandler(mockService)
	handler.RegisterMerchantHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status harus 200, dapat %d", rr.Code)
	}

	var response model.MerchantRegisterResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed parse response: %v", err)
	}

	if response.Message != "Registrasi merchant berhasil!" {
		t.Errorf("expected registrasi berhasil, got %s", response.Message)
	}
}