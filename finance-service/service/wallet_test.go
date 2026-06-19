package service_test

import (
	"testing"
	"finance-service/mocks"
	"finance-service/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTopUp_Unit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockWalletRepository(ctrl)

	// Ekspektasi: GetBalance mengembalikan float64, UpdateBalance menerima float64
	mockRepo.EXPECT().GetBalance("USER123").Return(50000.0, nil).AnyTimes()
	mockRepo.EXPECT().UpdateBalance("USER123", 100000.0).Return(nil).AnyTimes()

	// Inisialisasi dengan 2 parameter sesuai konstruktor aslimu
	svc := service.NewWalletService(mockRepo, "http://pricing-service")
	
	// Panggil fungsi yang benar: ProcessTopUp dengan 3 parameter
	newBalance, err := svc.ProcessTopUp("USER123", 50000.0, "")

	assert.NoError(t, err)
	assert.Equal(t, 100000.0, newBalance)
}