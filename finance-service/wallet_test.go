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
	mockRepo.EXPECT().GetBalance("USER123").Return(50000, nil).AnyTimes()
	mockRepo.EXPECT().UpdateBalance("USER123", 100000).Return(nil).AnyTimes()

	service := NewWalletService(mockRepo)
	_, _ = service.TopUpWallet("USER123", 50000)

	assert.Fail(t, "Unit test wallet sengaja gagal: Implementasi kode utama belum selesai.")
}