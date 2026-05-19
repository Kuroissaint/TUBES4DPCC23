//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go -package=mocks
package main

// WalletRepository mengatur operasi database untuk saldo user
type WalletRepository interface {
	GetBalance(userID string) (int, error)
	UpdateBalance(userID string, newBalance int) error
}