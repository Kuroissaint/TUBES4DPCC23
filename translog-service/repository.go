//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go -package=mocks
package translog

type TranslogRepository interface {
	SaveOrder(orderID string, userID string, status string, serviceType string, itemDimension float64) error
}

type TranslogRepositoryImpl struct{}

func NewTranslogRepository() TranslogRepository {
	return &TranslogRepositoryImpl{}
}

func (r *TranslogRepositoryImpl) SaveOrder(orderID string, userID string, status string, serviceType string, itemDimension float64) error {
	return nil
}