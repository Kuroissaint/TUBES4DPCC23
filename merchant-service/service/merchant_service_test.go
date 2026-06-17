package service

import (
	"errors"
	"testing"

	"merchant-service/model"
)

type MockMerchantRepository struct{}

func (m MockMerchantRepository) GetByNamaToko(namaToko string) (*model.Merchant, error) {
	if namaToko == "Toko Sudah Ada" {
		return &model.Merchant{
			ID:       1,
			NamaToko: namaToko,
			Kota:     "Bandung",
		}, nil
	}
	return nil, errors.New("merchant tidak ditemukan")
}

func (m MockMerchantRepository) RegisterMerchant(req model.MerchantRegisterRequest, hashedPassword string) (*model.Merchant, error) {
	return &model.Merchant{
		ID:       1,
		UserID:   1,
		NamaToko: req.NamaToko,
		Kota:     req.Kota,
	}, nil
}

func (m MockMerchantRepository) GetMerchantByID(id int) (*model.Merchant, error) {
	return &model.Merchant{
		ID:       id,
		NamaToko: "Warung Test",
		Kota:     "Bandung",
	}, nil
}

func (m MockMerchantRepository) AddMenuItem(merchantID int, req model.MenuItemRequest) (*model.MenuItem, error) {
	return &model.MenuItem{
		ID:         1,
		MerchantID: merchantID,
		Nama:       req.Nama,
		Harga:      req.Harga,
	}, nil
}

func (m MockMerchantRepository) GetMenuByMerchantID(merchantID int) ([]model.MenuItem, error) {
	return []model.MenuItem{
		{ID: 1, MerchantID: merchantID, Nama: "Nasi Goreng", Harga: 15000},
	}, nil
}

func (m MockMerchantRepository) UpdateMenuItem(id int, req model.MenuItemRequest) (*model.MenuItem, error) {
	return &model.MenuItem{
		ID:   id,
		Nama: req.Nama,
	}, nil
}

func (m MockMerchantRepository) DeleteMenuItem(id int) error {
	return nil
}

func TestRegisterMerchantSuccess(t *testing.T) {
	mockRepo := MockMerchantRepository{}
	merchantService := NewMerchantService(mockRepo)

	resp, err := merchantService.RegisterMerchant(model.MerchantRegisterRequest{
		Nama:     "Pak Budi",
		Email:    "pakbudi@gmail.com",
		NoHp:     "081234567890",
		Password: "password123",
		NamaToko: "Warung Pak Budi",
		Kota:     "Bandung",
	})

	if err != nil {
		t.Errorf("harusnya registrasi berhasil, tapi error: %v", err)
		return
	}

	if resp.UserID == 0 {
		t.Error("user ID tidak boleh kosong")
	}

	if resp.Message != "Registrasi merchant berhasil!" {
		t.Errorf("expected registrasi berhasil, got %s", resp.Message)
	}
}

func TestRegisterMerchantNamaTokoSudahAda(t *testing.T) {
	mockRepo := MockMerchantRepository{}
	merchantService := NewMerchantService(mockRepo)

	_, err := merchantService.RegisterMerchant(model.MerchantRegisterRequest{
		Nama:     "Pak Budi",
		Email:    "pakbudi2@gmail.com",
		Password: "password123",
		NamaToko: "Toko Sudah Ada",
	})

	if err == nil {
		t.Error("harusnya gagal karena nama toko sudah ada")
	}
}
