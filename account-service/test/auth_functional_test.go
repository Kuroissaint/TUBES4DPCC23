package test

import (
	"account-service/handler"
	"account-service/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAuthService struct{}

func (m MockAuthService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	return &model.LoginResponse{
		UserID: 1,
		Token: "dummy-token",
	}, nil
}

//driver
func (m MockAuthService) Register(req model.RegisterRequest) (*model.RegisterResponse, error) {
	return &model.RegisterResponse{
		UserID:  1,
		Message: "Registrasi berhasil!",
	}, nil
}

//driver
func (m MockAuthService) RegisterDriver(req model.DriverRegisterRequest) (*model.RegisterResponse, error) {
	return &model.RegisterResponse{
		UserID:  1,
		Message: "Registrasi driver berhasil!",
	}, nil
}

//cutomer
func (m MockAuthService) RegisterCustomer(req model.CustomerRegisterRequest) (*model.RegisterResponse, error) {
	return &model.RegisterResponse{
		UserID:  1,
		Message: "Registrasi customer berhasil!",
	}, nil
}


func TestLogin_Functional(t *testing.T) {

	mockService := MockAuthService{}
	authHandler := handler.NewAuthHandler(mockService)

	body := []byte(`{
		"email":"user@mail.com",
		"password":"123"
	}`)

	req := httptest.NewRequest(
		"POST",
		"/api/auth/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	authHandler.LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}

	var response model.LoginResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed parse response: %v", err)
	}

	if response.Token != "dummy-token" {
		t.Errorf("Expected dummy-token, got %s", response.Token)
	}

	if response.Token == "" {
		t.Errorf("Token tidak boleh kosong")
	}
}

