package repository

import (
	"database/sql"
	"fmt"
	"merchant-service/model"
)

type merchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) GetByNamaToko(namaToko string) (*model.Merchant, error) {
	merchant := &model.Merchant{}
	query := `SELECT id, user_id, nama_toko, kategori, kota, alamat, no_telp, is_open 
			  FROM merchants WHERE nama_toko = $1`
	err := r.db.QueryRow(query, namaToko).Scan(
		&merchant.ID,
		&merchant.UserID,
		&merchant.NamaToko,
		&merchant.Kategori,
		&merchant.Kota,
		&merchant.Alamat,
		&merchant.NoTelp,
		&merchant.IsOpen,
	)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

func (r *merchantRepository) RegisterMerchant(req model.MerchantRegisterRequest, hashedPassword string) (*model.Merchant, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	// 1. Insert ke tabel users
	var userID int
	err = tx.QueryRow(`
		INSERT INTO users (nama, email, no_hp, password, role)
		VALUES ($1, $2, $3, $4, 'merchant')
		RETURNING id`,
		req.Nama, req.Email, req.NoHp, hashedPassword,
	).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan user: %v", err)
	}

	// 2. Insert ke tabel merchants
	var merchantID int
	err = tx.QueryRow(`
		INSERT INTO merchants (user_id, nama_toko, kategori, kota, alamat, no_telp)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		userID, req.NamaToko, req.Kategori, req.Kota, req.Alamat, req.NoTelp,
	).Scan(&merchantID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal simpan merchant: %v", err)
	}

	tx.Commit()

	return &model.Merchant{
		ID:       merchantID,
		UserID:   userID,
		NamaToko: req.NamaToko,
		Kategori: req.Kategori,
		Kota:     req.Kota,
		Alamat:   req.Alamat,
		NoTelp:   req.NoTelp,
		IsOpen:   true,
	}, nil
}

func (r *merchantRepository) GetMerchantByID(id int) (*model.Merchant, error) {
	merchant := &model.Merchant{}
	query := `SELECT id, user_id, nama_toko, kategori, kota, alamat, no_telp, is_open
			  FROM merchants WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&merchant.ID,
		&merchant.UserID,
		&merchant.NamaToko,
		&merchant.Kategori,
		&merchant.Kota,
		&merchant.Alamat,
		&merchant.NoTelp,
		&merchant.IsOpen,
	)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

func (r *merchantRepository) AddMenuItem(merchantID int, req model.MenuItemRequest) (*model.MenuItem, error) {
	item := &model.MenuItem{}
	query := `INSERT INTO menu_items (merchant_id, nama, deskripsi, harga, kategori, is_tersedia)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, merchant_id, nama, deskripsi, harga, kategori, is_tersedia`
	err := r.db.QueryRow(query,
		merchantID, req.Nama, req.Deskripsi, req.Harga, req.Kategori, req.IsTersedia,
	).Scan(
		&item.ID,
		&item.MerchantID,
		&item.Nama,
		&item.Deskripsi,
		&item.Harga,
		&item.Kategori,
		&item.IsTersedia,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal tambah menu: %v", err)
	}
	return item, nil
}

func (r *merchantRepository) GetMenuByMerchantID(merchantID int) ([]model.MenuItem, error) {
	query := `SELECT id, merchant_id, nama, deskripsi, harga, kategori, is_tersedia
			  FROM menu_items WHERE merchant_id = $1`
	rows, err := r.db.Query(query, merchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.MenuItem
	for rows.Next() {
		var item model.MenuItem
		err := rows.Scan(
			&item.ID,
			&item.MerchantID,
			&item.Nama,
			&item.Deskripsi,
			&item.Harga,
			&item.Kategori,
			&item.IsTersedia,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *merchantRepository) UpdateMenuItem(id int, req model.MenuItemRequest) (*model.MenuItem, error) {
	item := &model.MenuItem{}
	query := `UPDATE menu_items SET nama=$1, deskripsi=$2, harga=$3, kategori=$4, is_tersedia=$5
			  WHERE id=$6
			  RETURNING id, merchant_id, nama, deskripsi, harga, kategori, is_tersedia`
	err := r.db.QueryRow(query,
		req.Nama, req.Deskripsi, req.Harga, req.Kategori, req.IsTersedia, id,
	).Scan(
		&item.ID,
		&item.MerchantID,
		&item.Nama,
		&item.Deskripsi,
		&item.Harga,
		&item.Kategori,
		&item.IsTersedia,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal update menu: %v", err)
	}
	return item, nil
}

func (r *merchantRepository) DeleteMenuItem(id int) error {
	_, err := r.db.Exec(`DELETE FROM menu_items WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("gagal hapus menu: %v", err)
	}
	return nil
}
