package service

import (
	"errors"
	"merchant-service/model"
	"merchant-service/repository"

	"golang.org/x/crypto/bcrypt"
)

type MerchantService interface {
	RegisterMerchant(req model.MerchantRegisterRequest) (*model.MerchantRegisterResponse, error)
	GetMerchant(id int) (*model.Merchant, error)
	AddMenuItem(merchantID int, req model.MenuItemRequest) (*model.MenuItemResponse, error)
	GetMenu(merchantID int) ([]model.MenuItem, error)
}	



type merchantService struct {
	repo repository.MerchantRepository
}

func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{repo: repo}
}

func (s *merchantService) RegisterMerchant(req model.MerchantRegisterRequest) (*model.MerchantRegisterResponse, error) {
	// Cek apakah nama toko sudah ada
	existing, _ := s.repo.GetByNamaToko(req.NamaToko)
	if existing != nil {
		return nil, errors.New("nama toko sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	// Simpan ke DB
	merchant, err := s.repo.RegisterMerchant(req, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &model.MerchantRegisterResponse{
		UserID:  merchant.UserID,
		Message: "Registrasi merchant berhasil!",
	}, nil
}

func (s *merchantService) GetMerchant(id int) (*model.Merchant, error) {
	merchant, err := s.repo.GetMerchantByID(id)
	if err != nil {
		return nil, errors.New("merchant tidak ditemukan")
	}
	return merchant, nil
}

func (s *merchantService) AddMenuItem(merchantID int, req model.MenuItemRequest) (*model.MenuItemResponse, error) {
	item, err := s.repo.AddMenuItem(merchantID, req)
	if err != nil {
		return nil, err
	}

	return &model.MenuItemResponse{
		ID:      item.ID,
		Message: "Menu berhasil ditambahkan!",
	}, nil
}

func (s *merchantService) GetMenu(merchantID int) ([]model.MenuItem, error) {
	return s.repo.GetMenuByMerchantID(merchantID)
}
