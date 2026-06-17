// menerima HTTP request/API register merchant
package handler

import (
	"encoding/json"
	"merchant-service/model"
	"merchant-service/service"
	"net/http"
	"strconv"
	"strings"
)

type MerchantHandler struct {
	merchantService service.MerchantService
}

func NewMerchantHandler(ms service.MerchantService) *MerchantHandler {
	return &MerchantHandler{merchantService: ms}
}

// REGISTER TOKO BARU
func (h *MerchantHandler) RegisterMerchantHandler(w http.ResponseWriter, r *http.Request) {
	var req model.MerchantRegisterRequest
	// Decode JSON body dari request ke struct MerchantRequest
	json.NewDecoder(r.Body).Decode(&req)

	// Validasi field wajib
	if req.Nama == "" || req.Email == "" || req.Password == "" || req.NamaToko == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "nama, email, password, dan nama_toko wajib diisi"})
		return
	}

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

// Lihat info merchant by ID
func (h *MerchantHandler) MerchantRouterHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// POST /api/merchant/12/menu → tambah menu
	if r.Method == "POST" && strings.HasSuffix(path, "/menu") {
		h.AddMenuItemHandler(w, r)
		return
	}

	// GET /api/merchant/12/menu → lihat menu
	if r.Method == "GET" && strings.HasSuffix(path, "/menu") {
		h.GetMenuHandler(w, r)
		return
	}

	// GET /api/merchant/12 → lihat info merchant
	if r.Method == "GET" {
		h.GetMerchantHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// Lihat info merchant by ID
func (h *MerchantHandler) GetMerchantHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/merchant/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id tidak valid"})
		return
	}

	resp, err := h.merchantService.GetMerchant(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "merchant tidak ditemukan"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Tambah menu item
func (h *MerchantHandler) AddMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil merchant ID dari URL: /api/merchant/12/menu
	// URL: /api/merchant/12/menu → split → ["", "api", "merchant", "12", "menu"]
	path := strings.TrimPrefix(r.URL.Path, "/api/merchant/")
	path = strings.TrimSuffix(path, "/menu")

	merchantID, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "merchant id tidak valid"})
		return
	}

	var req model.MenuItemRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Nama == "" || req.Harga == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "nama dan harga wajib diisi"})
		return
	}

	resp, err := h.merchantService.AddMenuItem(merchantID, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Lihat semua menu merchant
func (h *MerchantHandler) GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	merchantID, err := strconv.Atoi(parts[3])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "merchant id tidak valid"})
		return
	}

	resp, err := h.merchantService.GetMenu(merchantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
