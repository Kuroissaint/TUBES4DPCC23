// menerima HTTP request/API register merchant
package handler

import (
	"encoding/json"
	"merchant-service/model"
	"merchant-service/service"
	"net/http"
)

type MerchantHandler struct {
	merchantService service.MerchantService
}

func NewMerchantHandler(ms service.MerchantService) *MerchantHandler {
	return &MerchantHandler{merchantService: ms}
}

func (h *MerchantHandler) RegisterMerchantHandler(w http.ResponseWriter, r *http.Request) {
	var req model.MerchantRequest
	// Decode JSON body dari request ke struct MerchantRequest
	json.NewDecoder(r.Body).Decode(&req)

	// Panggil core bisnis logic di lapisan service
	resp, err := h.merchantService.RegisterMerchant(req)
	if err != nil {
		// Jika service mengembalikan error, kirim status 401 atau 400 sesuai kebutuhan bisnis
		w.WriteHeader(http.StatusBadRequest) // Diubah ke BadRequest karena skenario pendaftaran gagal
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Jika sukses, kirim respon JSON dengan status 200 OK
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}