package service

import (
	"account-service/model"
	"errors"
	"testing"
)

type MockUserRepository struct{}

func (m MockUserRepository) GetByEmail(email string) (*model.User, error) {

	if email == "firda@gmail.com" {
		return &model.User{
			ID:       1,
			Email:    email,
			Password: "123",
			Role:     "user",
		}, nil
	}

	return nil, errors.New("user tidak ditemukan")
}

func TestLoginSuccess(t *testing.T) {

	mockRepo := MockUserRepository{}

	authService := NewAuthService(mockRepo)

	resp, err := authService.Login(model.LoginRequest{
		Email:    "firda@gmail.com",
		Password: "123",
	})

	if err != nil {
		t.Error("harusnya login berhasil")
	}

	if resp.Token == "" {
		t.Error("token tidak boleh kosong")
	}
}