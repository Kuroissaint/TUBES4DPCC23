package service_test

import (
	"testing"
	"pricing-service/mocks" 
	"pricing-service/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalPrice_Unit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPromoRepository(ctrl)

	mockRepo.EXPECT().GetDiscountByCode("DISKON10").Return(10, nil).AnyTimes()

	svc := service.NewPricingService(mockRepo)
	
	finalPrice, err := svc.CalculateFinalPrice(100000, "DISKON10")

	assert.NoError(t, err)
	assert.Equal(t, int(90000), finalPrice, "Harga akhir harusnya 90000 setelah diskon 10%")
}
