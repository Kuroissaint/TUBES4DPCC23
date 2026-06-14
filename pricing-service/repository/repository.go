//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

type PromoRepository interface {
	GetDiscountByCode(code string) (int, error)
}
