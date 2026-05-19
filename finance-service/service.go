package main

type WalletService struct {
	repo WalletRepository
}

func NewWalletService(r WalletRepository) *WalletService {
	return &WalletService{repo: r}
}

// TopUpWallet adalah fungsi simulasi isi saldo
func (s *WalletService) TopUpWallet(userID string, amount int) (int, error) {
	// TODO: Nanti coding logika penambahan saldo di sini
	return 0, nil
}