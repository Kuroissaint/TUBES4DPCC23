package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dispatch-service/service"
)

func TestCreateOrderHandler_InvalidMethod(t *testing.T) {
	svc := service.NewDispatchService(nil)
	hdl := NewDispatchHandler(svc)

	// Skenario: Kirim GET padahal endpoint maunya POST
	req, err := http.NewRequest(http.MethodGet, "/api/dispatch/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.CreateOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestCreateOrderHandler_InvalidBody(t *testing.T) {
	svc := service.NewDispatchService(nil)
	hdl := NewDispatchHandler(svc)

	// Skenario: Kirim JSON rusak/kosong
	reqBody, _ := json.Marshal(map[string]string{"order_id": ""})
	req, err := http.NewRequest(http.MethodPost, "/api/dispatch/orders", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hdl.CreateOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid body: got %v want %v", status, http.StatusBadRequest)
	}
}