package model

// Data toko merchant (isi tabel merchants)
type Merchant struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	NamaToko  string `json:"nama_toko"`
	Kategori  string `json:"kategori"`
	Kota      string `json:"kota"`
	Alamat    string `json:"alamat"`
	NoTelp    string `json:"no_telp"`
	IsOpen    bool   `json:"is_open"`
	CreatedAt string `json:"created_at"`
}

// Request saat merchant register (isi tabel users + merchants sekaligus)
type MerchantRegisterRequest struct {
	// Data untuk tabel users
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
	// Data untuk tabel merchants
	NamaToko string `json:"nama_toko"`
	Kategori string `json:"kategori"`
	Kota     string `json:"kota"`
	Alamat   string `json:"alamat"`
	NoTelp   string `json:"no_telp"`
}

// Response setelah merchant register berhasil
type MerchantRegisterResponse struct {
	UserID   int    `json:"user_id"`
	Message  string `json:"message"`
}

// Data menu item (isi tabel menu_items)
type MenuItem struct {
	ID         int    `json:"id"`
	MerchantID int    `json:"merchant_id"`
	Nama       string `json:"nama"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga"`
	Kategori   string `json:"kategori"`
	IsTersedia bool   `json:"is_tersedia"`
}

// Request saat tambah/update menu
type MenuItemRequest struct {
	Nama       string `json:"nama"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga"`
	Kategori   string `json:"kategori"`
	IsTersedia bool   `json:"is_tersedia"`
}

// Response setelah tambah/update menu
type MenuItemResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
