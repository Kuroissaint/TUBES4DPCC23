package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"location-service/service"
)

func TestUpdateLocationHandler(t *testing.T) {
	// Setup service asli dengan parameter nil (mock database)
	svc := service.NewLocationService(nil)
	hdl := NewLocationHandler(svc)

	// Skenario 1: Test Request POST Valid
	reqBody, _ := json.Marshal(UpdateLocationRequest{
		DriverID:  "DRV-TEST-99",
		Latitude:  -6.2000,
		Longitude: 106.8166,
	})

	req, err := http.NewRequest(http.MethodPost, "/api/location/update", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.UpdateLocationHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Skenario 2: Test Request Method Salah (Harus POST, kita kirim GET)
	reqGet, _ := http.NewRequest(http.MethodGet, "/api/location/update", nil)
	rrGet := httptest.NewRecorder()
	handler.ServeHTTP(rrGet, reqGet)

	if status := rrGet.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for GET: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestGetNearbyDriversHandler(t *testing.T) {
	svc := service.NewLocationService(nil)
	hdl := NewLocationHandler(svc)

	// Skenario: Ambil driver terdekat dari query param URL
	req, err := http.NewRequest(http.MethodGet, "/api/location/nearby?lat=-6.2000&lon=106.8166", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.GetNearbyDriversHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["status"] != "success" {
		t.Errorf("expected success status, got %v", response["status"])
	}
}