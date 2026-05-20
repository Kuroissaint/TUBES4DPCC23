package main

import (
	"testing"
	"pricing-service/mocks" 
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalPrice_Unit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPromoRepository(ctrl)

	// Ekspektasi: Kode DISKON10 akan memberikan potongan 10%
	mockRepo.EXPECT().GetDiscountByCode("DISKON10").Return(10, nil).AnyTimes()

	service := NewPricingService(mockRepo)
	
	// Hitung harga akhir: 100000 dipotong 10%
	finalPrice, err := service.CalculateFinalPrice(100000, "DISKON10")

	assert.NoError(t, err)
	// Pastikan menggunakan int(90000) agar seragam dengan return fungsi
	assert.Equal(t, int(90000), finalPrice, "Harga akhir harusnya 90000 setelah diskon 10%")
}