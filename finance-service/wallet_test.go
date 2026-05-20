package main

import (
	"testing"
	"finance-service/mocks" // Sesuaikan dengan nama module di go.mod kamu
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

	service := NewWalletService(mockRepo)
	
	// Eksekusi TopUp sebesar 50.000
	newBalance, err := service.TopUpWallet("USER123", 50000)

	// Verifikasi hasil evaluasi Unit Test
	assert.NoError(t, err)
	assert.Equal(t, int(100000), newBalance, "Saldo akhir harusnya 100000 setelah topup 50000")
}