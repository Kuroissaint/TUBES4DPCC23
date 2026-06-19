package repository

import (
	"database/sql"
	"errors"
)

type WalletRepository interface {
	GetBalance(userID string) (float64, error)
	UpdateBalance(userID string, newBalance float64) error
}

type sqlWalletRepo struct {
	db *sql.DB
}

func NewSqlWalletRepo(db *sql.DB) WalletRepository {
	return &sqlWalletRepo{db: db}
}

func (r *sqlWalletRepo) GetBalance(userID string) (float64, error) {
	var balance float64
	query := "SELECT balance FROM wallets WHERE user_id = $1"
	
	err := r.db.QueryRow(query, userID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user tidak ditemukan")
		}
		return 0, err
	}
	
	return balance, nil
}

func (r *sqlWalletRepo) UpdateBalance(userID string, newBalance float64) error {
	query := "UPDATE wallets SET balance = $1 WHERE user_id = $2"
	_, err := r.db.Exec(query, newBalance, userID)
	return err
}