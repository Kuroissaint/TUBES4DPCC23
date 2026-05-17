package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"merchant-service/handler"
	"merchant-service/model"
)

// MockMerchantService digunakan untuk memotong dependency database asli di level functional test lokal
type MockMerchantService struct{}

func (m MockMerchantService) RegisterMerchant(req model.MerchantRequest) (*model.MerchantResponse, error) {
	return &model.MerchantResponse{
		MerchantID: 99,
		Status:     "active",
	}, nil
}

func TestRegisterMerchant_Functional(t *testing.T) {
	// 1. Inisialisasi mock service dan inject ke handler utama merchant
	mockService := MockMerchantService{}
	merchantHandler := handler.NewMerchantHandler(mockService)

	// 2. Siapkan payload JSON request body
	body := []byte(`{
		"name": "Geprek Bensu UPI",
		"city": "Bandung"
	}`)

	// 3. Buat request HTTP palsu yang menembak route register merchant
	req := httptest.NewRequest(
		"POST",
		"/api/merchants/register",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	// 4. Siapkan ResponseRecorder untuk menampung hasil eksekusi
	rr := httptest.NewRecorder()

	// 5. Jalankan HTTP handler
	merchantHandler.RegisterMerchantHandler(rr, req)

	// 6. Verifikasi HTTP Status Code (Harus 200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}

	// 7. Parsing response JSON ke dalam struct model.MerchantResponse
	var response model.MerchantResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed parse response: %v", err)
	}

	// 8. Validasi isi field response body
	if response.MerchantID != 99 {
		t.Errorf("Expected MerchantID 99, got %d", response.MerchantID)
	}

	if response.Status != "active" {
		t.Errorf("Expected status active, got %s", response.Status)
	}
}