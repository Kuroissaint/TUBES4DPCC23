package handler

import (
	"account-service/model"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

type MockAuthService struct{}

func (m MockAuthService) Login(req model.LoginRequest) (*model.LoginResponse, error) {

	return &model.LoginResponse{
		Token: "dummy-token",
	}, nil
}

func TestLoginHandler(t *testing.T) {

	jsonBody := []byte(`{
		"email":"firda@gmail.com",
		"password":"123"
	}`)

	req := httptest.NewRequest(
		"POST",
		"/login",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mockService := MockAuthService{}
	handler := NewAuthHandler(mockService)

	//Menjalankan hendler
	handler.LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("status harus 200")
	}

	// parsing response body
	var response model.LoginResponse

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed parse response: %v", err)
	}

	// cek token
	if response.Token != "dummy-token" {
		t.Errorf("expected dummy-token, got %s", response.Token)
	}
}
