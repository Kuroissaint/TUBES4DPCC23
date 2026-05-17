package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"merchant-service/model"
)

// MockMerchantService digunakan untuk menyimulasikan core business logic
type MockMerchantService struct{}

func (m MockMerchantService) RegisterMerchant(req model.MerchantRequest) (*model.MerchantResponse, error) {
	return &model.MerchantResponse{
		MerchantID: 99,
		Status:     "active",
	}, nil
}

func TestRegisterMerchantHandler(t *testing.T) {
	// 1. Siapkan mock request body JSON
	jsonBody := []byte(`{
		"name": "Geprek Bensu UPI",
		"city": "Bandung"
	}`)

	// 2. Buat objek request HTTP palsu menembak ke endpoint merchant
	req := httptest.NewRequest(
		"POST",
		"/merchants",
		bytes.NewBuffer(jsonBody),
	)
	req.Header.Set("Content-Type", "application/json")

	// 3. Buat ResponseRecorder untuk menangkap output dari handler
	rr := httptest.NewRecorder()

	// 4. Inisialisasi handler dengan dependensi mock service
	mockService := MockMerchantService{}
	handler := NewMerchantHandler(mockService) // Pastikan fungsi ini didefinisikan di merchant_handler.go

	// 5. Jalankan handler
	handler.RegisterMerchantHandler(rr, req)

	// 6. Verifikasi HTTP Status Code (Harus 200 OK atau 201 Created tergantung rancanganmu)
	if rr.Code != http.StatusOK {
		t.Error("status harus 200")
	}

	// 7. Parsing response body dari JSON kembali ke struct model
	var response model.MerchantResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed parse response: %v", err)
	}

	// 8. Cek validitas data hasil mock di dalam response body
	if response.MerchantID != 99 {
		t.Errorf("expected MerchantID 99, got %d", response.MerchantID)
	}

	if response.Status != "active" {
		t.Errorf("expected status active, got %s", response.Status)
	}
}