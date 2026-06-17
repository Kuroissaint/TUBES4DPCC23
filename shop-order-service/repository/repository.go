//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq" // Butuh library pq buat handle array string
)

type ShopOrderRepository interface {
	SaveCart(orderID string, userID string, merchantID string, items []string, status string) error
}

type ShopOrderRepositoryImpl struct {
	DB *sql.DB
}

func NewShopOrderRepository(db *sql.DB) ShopOrderRepository {
	return &ShopOrderRepositoryImpl{DB: db}
}

func (r *ShopOrderRepositoryImpl) SaveCart(orderID string, userID string, merchantID string, items []string, status string) error {
	// Simpan array "items" langsung pake pq.Array(items)
	query := `INSERT INTO shop_orders (order_id, user_id, merchant_id, items, status) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(query, orderID, userID, merchantID, pq.Array(items), status)
	if err != nil {
		return fmt.Errorf("failed to insert shop order data: %v", err)
	}

	return nil
}