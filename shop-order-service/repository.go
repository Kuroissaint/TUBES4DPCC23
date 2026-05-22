//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go -package=mocks
package shoporder

type ShopOrderRepository interface {
	SaveCart(orderID string, userID string, merchantID string, items []string, status string) error
}

type ShopOrderRepositoryImpl struct{}

func NewShopOrderRepository() ShopOrderRepository {
	return &ShopOrderRepositoryImpl{}
}

func (r *ShopOrderRepositoryImpl) SaveCart(orderID string, userID string, merchantID string, items []string, status string) error {
	return nil
}