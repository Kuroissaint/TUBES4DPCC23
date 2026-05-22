package translog

import (
	"fmt"
	"github.com/google/uuid"
)

type TransportOrder struct {
	OrderID       string  `json:"order_id"`
	UserID        string  `json:"user_id"`
	Status        string  `json:"status"`
	ServiceType   string  `json:"service_type"`
	ItemDimension float64 `json:"item_dimension,omitempty"`
}

type TranslogService interface {
	CreateTransportOrder() (*TransportOrder, error)
	ValidateStatusTransition(currentStatus, newStatus string) error
}

type TranslogServiceImpl struct {
	Repo TranslogRepository
}

func NewTranslogService(repo TranslogRepository) TranslogService {
	return &TranslogServiceImpl{Repo: repo}
}

func (s *TranslogServiceImpl) ValidateStatusTransition(currentStatus, newStatus string) error {
	if currentStatus == "SEARCHING" && newStatus == "COMPLETED" {
		return fmt.Errorf("pesanan tidak bisa langsung COMPLETED dari SEARCHING")
	}
	return nil
}

func (s *TranslogServiceImpl) CreateTransportOrder() (*TransportOrder, error) {
	newOrder := &TransportOrder{
		OrderID:       uuid.New().String(),
		UserID:        uuid.New().String(),
		Status:        "SEARCHING",
		ServiceType:   "ride",
		ItemDimension: 0.0,
	}

	err := s.Repo.SaveOrder(newOrder.OrderID, newOrder.UserID, newOrder.Status, newOrder.ServiceType, newOrder.ItemDimension)
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}