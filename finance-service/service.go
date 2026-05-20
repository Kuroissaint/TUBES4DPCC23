package main

type WalletService struct {
	repo WalletRepository
}

func NewWalletService(r WalletRepository) *WalletService {
	return &WalletService{repo: r}
}

// TopUpWallet adalah fungsi simulasi isi saldo
func (s *WalletService) TopUpWallet(userID string, amount int) (int, error) {
	// 1. Cek saldo awal user dari database/mock
	currentBalance, err := s.repo.GetBalance(userID)
	if err != nil {
		return 0, err
	}

	// 2. Tambahkan saldo lama dengan jumlah top-up
	newBalance := currentBalance + amount

	// 3. Simpan perubahan saldo ke database/mock
	err = s.repo.UpdateBalance(userID, newBalance)
	if err != nil {
		return 0, err
	}

	// 4. Kembalikan saldo terbaru
	return newBalance, nil
}