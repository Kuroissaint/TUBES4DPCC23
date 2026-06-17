package service

import (
	"account-service/model"
	"errors"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct{}

func (m MockUserRepository) GetByEmail(email string) (*model.User, error) {
	if email == "firda@gmail.com" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

		return &model.User{
			ID:             1,
			Nama:           "Firda",      // ← tambah
			Email:          email,
			NoHp:           "081234567890", // ← tambah
			HashedPassword: string(hashedPassword),
			Role:           "user",
		}, nil
	}
	return nil, errors.New("user tidak ditemukan")
}

func (m MockUserRepository) Register(user model.User) (*model.User, error) {
	return &model.User{
		ID:    1,
		Nama:  user.Nama,
		Email: user.Email,
		NoHp:  user.NoHp,
		Role:  user.Role,
	}, nil
}

//mock driver
func (m MockUserRepository) RegisterDriver(user model.User, driver model.DriverProfile) (*model.User, error) {
	return &model.User{
		ID:    1,
		Nama:  user.Nama,
		Email: user.Email,
		NoHp:  user.NoHp,
		Role:  "driver",
	}, nil
}

//mock customer
func (m MockUserRepository) RegisterCustomer(user model.User, alamat string) (*model.User, error) {
	return &model.User{
		ID:    1,
		Nama:  user.Nama,
		Email: user.Email,
		NoHp:  user.NoHp,
		Role:  "user",
	}, nil
}


func TestLoginSuccess(t *testing.T) {
	mockRepo := MockUserRepository{}
	authService := NewAuthService(mockRepo)

	resp, err := authService.Login(model.LoginRequest{
		Email:    "firda@gmail.com",
		Password: "password123",
	})

	if err != nil {
		t.Errorf("harusnya login berhasil, tapi error: %v", err)
		return // ← tambah return supaya tidak panic
	}

	if resp == nil || resp.Token == "" {
		t.Error("token tidak boleh kosong")
	}
}

func TestLoginGagalPasswordSalah(t *testing.T) {
	mockRepo := MockUserRepository{}
	authService := NewAuthService(mockRepo)

	_, err := authService.Login(model.LoginRequest{
		Email:    "firda@gmail.com",
		Password: "passwordSalah",
	})

	if err == nil {
		t.Error("harusnya login gagal karena password salah")
	}
}

func TestLoginGagalEmailTidakAda(t *testing.T) {
	mockRepo := MockUserRepository{}
	authService := NewAuthService(mockRepo)

	_, err := authService.Login(model.LoginRequest{
		Email:    "tidakada@gmail.com",
		Password: "password123",
	})

	if err == nil {
		t.Error("harusnya login gagal karena email tidak ada")
	}
}

