package model

type User struct {
	ID             int    `json:"id"`
	Nama           string `json:"nama"`
	Email          string `json:"email"`
	NoHp           string `json:"no_hp"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password"`
	Role           string `json:"role"`
	CreatedAt      string `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// RegisterResponse adalah struktur respons dari backend setelah user berhasil mendaftar.
type RegisterResponse struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

//untuk memasukan datanya ke database driver//
//DRIVER//
type DriverRegisterRequest struct {
	Nama           string `json:"nama"`
	Email          string `json:"email"`
	NoHp           string `json:"no_hp"`
	Password       string `json:"password"`
	NoSim          string `json:"no_sim"`
	NoPlat         string `json:"no_plat"`
	JenisKendaraan string `json:"jenis_kendaraan"`
}
type DriverProfile struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	NoSim          string `json:"no_sim"`
	NoPlat         string `json:"no_plat"`
	JenisKendaraan string `json:"jenis_kendaraan"`
}

//CUSTOMER//
type CustomerRegisterRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
	Alamat   string `json:"alamat"`
}

type CustomerProfile struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Alamat string `json:"alamat"`
}
