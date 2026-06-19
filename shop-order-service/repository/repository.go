//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"database/sql"
	"fmt"
	"shop-order-service/model"

	"github.com/lib/pq"
)

type ShopOrderRepository interface {
	SaveCart(cart *model.ShoppingCart) error
}

type ShopOrderRepositoryImpl struct {
	DB *sql.DB
}

func NewShopOrderRepository(db *sql.DB) ShopOrderRepository {
	return &ShopOrderRepositoryImpl{DB: db}
}

func (r *ShopOrderRepositoryImpl) SaveCart(cart *model.ShoppingCart) error {
	query := `INSERT INTO shop_orders (order_id, user_id, merchant_id, items, total_price, delivery_address, payment_status, status) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.DB.Exec(query, cart.OrderID, cart.UserID, cart.MerchantID, pq.Array(cart.Items), cart.TotalPrice, cart.DeliveryAddress, cart.PaymentStatus, cart.Status)
	if err != nil {
		return fmt.Errorf("failed to insert shop order data: %v", err)
	}

	return nil
}