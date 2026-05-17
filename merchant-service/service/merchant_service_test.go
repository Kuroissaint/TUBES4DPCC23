package service

import (
	"errors"
	"testing"

	"merchant-service/model"
)

// MockMerchantRepository digunakan untuk menyimulasikan database lokal/memory
type MockMerchantRepository struct{}

func (m MockMerchantRepository) GetByName(name string) (*model.Merchant, error) {
	// Menyimulasikan bahwa merchant dengan nama ini SUDAH ADA di database
	if name == "Toko Sudah Ada" {
		return &model.Merchant{
			ID:   1,
			Name: name,
			City: "Bandung",
		}, nil
	}

	// Menyimulasikan merchant belum terdaftar (bisa didaftarkan)
	return nil, errors.New("merchant tidak ditemukan")
}

func TestRegisterMerchantSuccess(t *testing.T) {
	// 1. Inisialisasi mock repository
	mockRepo := MockMerchantRepository{}

	// 2. Inisialisasi service dengan menyuntikkan mock repository
	merchantService := NewMerchantService(mockRepo)

	// 3. Jalankan fungsi yang ingin diuji dengan data merchant baru
	resp, err := merchantService.RegisterMerchant(model.MerchantRequest{
		Name: "Geprek Bensu UPI",
		City: "Bandung",
	})

	// 4. Verifikasi hasil akhir (Assertion)
	if err != nil {
		t.Error("harusnya registrasi merchant berhasil")
	}

	if resp.MerchantID == 0 {
		t.Error("merchant ID tidak boleh kosong atau nol")
	}

	if resp.Status != "active" {
		t.Errorf("expected status active, got %s", resp.Status)
	}
}