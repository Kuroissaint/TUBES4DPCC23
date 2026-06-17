//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"database/sql"
	"fmt"
)

type TranslogRepository interface {
	SaveOrder(orderID string, userID string, status string, serviceType string, itemDimension float64) error
}

type TranslogRepositoryImpl struct {
	DB *sql.DB
}

// Fungsi New menerima objek database
func NewTranslogRepository(db *sql.DB) TranslogRepository {
	return &TranslogRepositoryImpl{DB: db}
}

func (r *TranslogRepositoryImpl) SaveOrder(orderID string, userID string, status string, serviceType string, itemDimension float64) error {
	// Query SQL beneran untuk simpan ke Supabase
	query := `INSERT INTO translogs (order_id, user_id, status, service_type, item_dimension) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(query, orderID, userID, status, serviceType, itemDimension)
	if err != nil {
		return fmt.Errorf("failed to insert translog data: %v", err)
	}
	
	return nil
}