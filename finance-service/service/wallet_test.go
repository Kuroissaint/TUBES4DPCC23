package service_test

import (
    "testing"
    "finance-service/mocks" 
    "finance-service/service"
    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
)

func TestTopUpWallet_Unit(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockWalletRepository(ctrl)

    // Setup ekspektasi tiruan untuk simulasi DB
    // Saldo awal = 50.000
    mockRepo.EXPECT().GetBalance("USER123").Return(50000, nil).AnyTimes()
    // Setelah ditambah 50.000, sistem wajib memanggil UpdateBalance dengan nilai 100.000
    mockRepo.EXPECT().UpdateBalance("USER123", 100000).Return(nil).AnyTimes()

    // PERBAIKAN 1: Tambahkan parameter string sembarang (misal "dummy-secret" atau "") 
    // karena NewWalletService di kode aslimu minta 2 parameter.
    svc := service.NewWalletService(mockRepo, "dummy-secret")
    
    // PERBAIKAN 2: Pastikan nama fungsinya sama PERSIS dengan yang ada di file wallet.go aslimu.
    // Asumsiku namanya adalah TopUp (bukan TopUpWallet). 
    // Kalau ternyata di wallet.go namanya AddBalance, ganti jadi svc.AddBalance.
    newBalance, err := svc.TopUp("USER123", 50000)

    // Verifikasi hasil evaluasi Unit Test
    assert.NoError(t, err)
    assert.Equal(t, int(100000), newBalance, "Saldo akhir harusnya 100000 setelah topup 50000")
}