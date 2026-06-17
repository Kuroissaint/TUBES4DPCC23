package repository

import (
	"account-service/model"
	"database/sql"
	"fmt" //memformat teks dan menampilkan data ke layar terminal (seperti log pesan)
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}

	query := `SELECT id, nama, email, no_hp, password, role 
			  FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Nama,
		&user.Email,
		&user.NoHp,
		&user.HashedPassword,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

//register function
func (r *userRepository) Register(user model.User) (*model.User, error) {
	query := `INSERT INTO users (nama, email, no_hp, password, role) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id, nama, email, no_hp, role`

	result := &model.User{}
	err := r.db.QueryRow(query,
		user.Nama,
		user.Email,
		user.NoHp,
		user.HashedPassword,
		user.Role,
	).Scan(
		&result.ID,
		&result.Nama,
		&result.Email,
		&result.NoHp,
		&result.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal register: %v", err)
	}

	return result, nil
}

//REGISTER DRIVER//
func (r *userRepository) RegisterDriver(user model.User, driver model.DriverProfile) (*model.User, error) {
	// Mulai transaction — kalau salah satu gagal, keduanya dibatalkan
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	// 1. Insert ke tabel users
	var userID int
	err = tx.QueryRow(`
		INSERT INTO users (nama, email, no_hp, password, role)
		VALUES ($1, $2, $3, $4, 'driver')
		RETURNING id`,
		user.Nama, user.Email, user.NoHp, user.HashedPassword,
	).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan user: %v", err)
	}

	// 2. Insert ke tabel driver_profiles
	_, err = tx.Exec(`
		INSERT INTO driver_profiles (user_id, no_sim, no_plat, jenis_kendaraan)
		VALUES ($1, $2, $3, $4)`,
		userID, driver.NoSim, driver.NoPlat, driver.JenisKendaraan,
	)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan driver profile: %v", err)
	}

	// Commit — semua berhasil
	tx.Commit()

	return &model.User{
		ID:    userID,
		Nama:  user.Nama,
		Email: user.Email,
		NoHp:  user.NoHp,
		Role:  "driver",
	}, nil
}

//Register CUSTOMER//
func (r *userRepository) RegisterCustomer(user model.User, alamat string) (*model.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	// 1. Insert ke tabel users
	var userID int
	err = tx.QueryRow(`
		INSERT INTO users (nama, email, no_hp, password, role)
		VALUES ($1, $2, $3, $4, 'user')
		RETURNING id`,
		user.Nama, user.Email, user.NoHp, user.HashedPassword,
	).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan user: %v", err)
	}

	// 2. Insert ke tabel customer_profiles
	_, err = tx.Exec(`
		INSERT INTO customer_profiles (user_id, alamat)
		VALUES ($1, $2)`,
		userID, alamat,
	)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan customer profile: %v", err)
	}

	tx.Commit()

	return &model.User{
		ID:    userID,
		Nama:  user.Nama,
		Email: user.Email,
		NoHp:  user.NoHp,
		Role:  "user",
	}, nil
}
