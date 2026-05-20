//menerima HTTP request/API login register
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