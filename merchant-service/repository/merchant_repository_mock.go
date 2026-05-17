package repository

import (
	"merchant-service/model"
	"database/sql"
)

type merchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) GetByName(name string) (*model.Merchant, error) {

	// sementara dummy dulu supaya gampang test
	return &model.Merchant{
		ID:       1,
		Name:     "Merchant Toko A",
		City: "Bandung",
	}, nil
}