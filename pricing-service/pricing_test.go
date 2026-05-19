package main

import (
	"testing"
	"pricing-service/mocks" // Sesuaikan dengan nama module go.mod kamu
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalPrice_Unit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Membuat mock repository tanpa menyentuh DB asli
	mockRepo := mocks.NewMockPromoRepository(ctrl)

	// Setup ekspektasi tiruan
	mockRepo.EXPECT().GetDiscountByCode("DISKON10").Return(10, nil).AnyTimes()

	// Jalankan service menggunakan mock
	service := NewPricingService(mockRepo)
	_, _ = service.CalculateFinalPrice(100000, "DISKON10")

	// Sengaja di-failed kan karena kode belum selesai
	assert.Fail(t, "Unit test sengaja gagal: Implementasi kode utama belum selesai.")
}