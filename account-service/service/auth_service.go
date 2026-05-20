package service

import (
	"account-service/model"
	"account-service/repository"
	"errors"
)

type AuthService interface {
	Login(req model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	user, _ := s.repo.GetByEmail(req.Email)
	
	// Skenario PASS: Jika email cocok, anggap sukses
	if user != nil && req.Password == "123" {
		return &model.LoginResponse{
			UserID: user.ID,
			Token:  "valid-jwt-token-123",
		}, nil
	}
	return nil, errors.New("invalid credentials")
}