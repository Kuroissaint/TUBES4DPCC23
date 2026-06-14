package service

import (
	"shop-order-service/model"
	"shop-order-service/repository"

	"github.com/google/uuid"
)

type ShopOrderService interface {
	CreateShoppingOrder() (*model.ShoppingCart, error)
}

type ShopOrderServiceImpl struct {
	Repo repository.ShopOrderRepository
}

func NewShopOrderService(repo repository.ShopOrderRepository) ShopOrderService {
	return &ShopOrderServiceImpl{Repo: repo}
}

func (s *ShopOrderServiceImpl) CreateShoppingOrder() (*model.ShoppingCart, error) {
	cart := &model.ShoppingCart{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MerchantID: uuid.New().String(),
		Items:      []string{"Kopi Gula Aren", "Roti Bakar"},
		Status:     "AT_STORE",
	}

	err := s.Repo.SaveCart(cart.OrderID, cart.UserID, cart.MerchantID, cart.Items, cart.Status)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
