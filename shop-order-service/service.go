package shoporder

import "github.com/google/uuid"

type ShoppingCart struct {
	OrderID    string
	UserID     string
	MerchantID string
	Items      []string
	Status     string
}

func (c *ShoppingCart) AddToCart(item string) {
	c.Items = append(c.Items, item)
}

type ShopOrderService interface {
	CreateShoppingOrder() (*ShoppingCart, error)
}

type ShopOrderServiceImpl struct {
	Repo ShopOrderRepository
}

func NewShopOrderService(repo ShopOrderRepository) ShopOrderService {
	return &ShopOrderServiceImpl{Repo: repo}
}

func (s *ShopOrderServiceImpl) CreateShoppingOrder() (*ShoppingCart, error) {
	cart := &ShoppingCart{
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