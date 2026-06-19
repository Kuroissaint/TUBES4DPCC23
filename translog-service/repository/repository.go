//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"database/sql"
	"fmt"
	"translog-service/model"
)

type TranslogRepository interface {
	SaveOrder(order *model.TransportOrder) error
}

type TranslogRepositoryImpl struct {
	DB *sql.DB
}

func NewTranslogRepository(db *sql.DB) TranslogRepository {
	return &TranslogRepositoryImpl{DB: db}
}

func (r *TranslogRepositoryImpl) SaveOrder(order *model.TransportOrder) error {
	query := `INSERT INTO translogs (order_id, user_id, driver_id, service_type, pickup_location, dropoff_location, item_dimension, fee, status) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.DB.Exec(query, order.OrderID, order.UserID, order.DriverID, order.ServiceType, order.PickupLocation, order.DropoffLocation, order.ItemDimension, order.Fee, order.Status)
	if err != nil {
		return fmt.Errorf("failed to insert translog data: %v", err)
	}

	return nil
}