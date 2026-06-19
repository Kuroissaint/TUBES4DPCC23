//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"database/sql"
	"errors"
)

type PromoRepository interface {
	GetDiscountByCode(code string) (int, error)
}

type sqlPromoRepo struct {
	db *sql.DB
}

func NewSqlPromoRepo(db *sql.DB) PromoRepository {
	return &sqlPromoRepo{db: db}
}

func (r *sqlPromoRepo) GetDiscountByCode(code string) (int, error) {
	var discount int
	var isActive bool

	query := "SELECT discount_percentage, is_active FROM promos WHERE code = $1"
	err := r.db.QueryRow(query, code).Scan(&discount, &isActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("promo tidak ditemukan")
		}
		return 0, err
	}

	if !isActive {
		return 0, errors.New("promo sudah tidak aktif")
	}

	return discount, nil
}