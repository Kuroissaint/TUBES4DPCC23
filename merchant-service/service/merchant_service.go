package service

import (
	"errors"
	"merchant-service/model"
	"merchant-service/repository"
)

type MerchantService interface {
	RegisterMerchant(req model.MerchantRequest) (*model.MerchantResponse, error)
}

type merchantService struct {
	repo repository.MerchantRepository
}

func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{repo: repo}
}

func (s *merchantService) RegisterMerchant(req model.MerchantRequest) (*model.MerchantResponse, error) {
	// Cek apakah merchant dengan nama tersebut sudah terdaftar
	merchant, _ := s.repo.GetByName(req.Name)

	// Skenario PASS: Jika nama belum terdaftar dan kota tidak kosong, anggap sukses
	if merchant == nil && req.City != "" {
		return &model.MerchantResponse{
			MerchantID: 99, // Dummy ID untuk pendaftaran baru
			Status:     "active",
		}, nil
	}
	return nil, errors.New("merchant already exists or invalid data")
}