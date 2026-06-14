package service

import (
	"fmt"
	"github.com/google/uuid"
	"translog-service/model"
	"translog-service/repository"
)

type TranslogService interface {
	CreateTransportOrder() (*model.TransportOrder, error)
	ValidateStatusTransition(currentStatus, newStatus string) error
}

type TranslogServiceImpl struct {
	Repo repository.TranslogRepository
}

func NewTranslogService(repo repository.TranslogRepository) TranslogService {
	return &TranslogServiceImpl{Repo: repo}
}

func (s *TranslogServiceImpl) ValidateStatusTransition(currentStatus, newStatus string) error {
	if currentStatus == "SEARCHING" && newStatus == "COMPLETED" {
		return fmt.Errorf("pesanan tidak bisa langsung COMPLETED dari SEARCHING")
	}
	return nil
}

func (s *TranslogServiceImpl) CreateTransportOrder() (*model.TransportOrder, error) {
	newOrder := &model.TransportOrder{
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
