// menerima HTTP request/API login register
package handler

import (
	"account-service/model"
	"account-service/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(as service.AuthService) *AuthHandler {
	return &AuthHandler{authService: as}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp, err := h.authService.Login(req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(resp)
}

// register function - sekarang hanya kasih pesan arah
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "gunakan endpoint yang sesuai: /api/customer/register, /api/driver/register, atau /api/merchant/register",
	})
}

// register customer
func (h *AuthHandler) CustomerRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CustomerRegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Nama == "" || req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "nama, email, dan password wajib diisi"})
		return
	}

	if len(req.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "password minimal 8 karakter"})
		return
	}

	resp, err := h.authService.RegisterCustomer(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(resp)
}

// register driver
func (h *AuthHandler) DriverRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req model.DriverRegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Nama == "" || req.Email == "" || req.Password == "" || req.NoSim == "" || req.NoPlat == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "semua field wajib diisi"})
		return
	}

	if len(req.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "password minimal 8 karakter"})
		return
	}

	resp, err := h.authService.RegisterDriver(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(resp)
}

