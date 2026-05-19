//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go -package=mocks
package main

// PromoRepository adalah kontrak fungsi database
type PromoRepository interface {
	GetDiscountByCode(code string) (int, error)
}