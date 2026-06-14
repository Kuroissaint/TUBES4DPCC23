package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/mocks"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/model"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/service"
)

func TestSubmitReview_Failed_RatingTooHigh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	svc := service.NewReviewService(mockRepo)

	invalidReview := model.Review{
		OrderID:  "ORD-111",
		UserID:   "USR-222",
		DriverID: "DRV-333",
		Rating:   6, 
		Comment:  "Driver ngebut banget",
	}

	mockRepo.EXPECT().SaveReview(gomock.Any(), gomock.Any()).Times(0)

	err := svc.SubmitReview(context.Background(), invalidReview)

	if err == nil {
		t.Errorf("Ekspektasi muncul error validasi rating, tapi malah sukses")
	}
	expectedErr := "rating tidak valid: harus berada di antara rentang angka 1 sampai 5"
	if err.Error() != expectedErr {
		t.Errorf("Ekspektasi pesan error '%s', mendapat: '%v'", expectedErr, err.Error())
	}
}

func TestSubmitReview_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	svc := service.NewReviewService(mockRepo)

	validReview := model.Review{
		OrderID:  "ORD-111",
		UserID:   "USR-222",
		DriverID: "DRV-333",
		Rating:   5, 
		Comment:  "Pelayanan sangat ramah dan memuaskan!",
	}

	mockRepo.EXPECT().SaveReview(gomock.Any(), validReview).Return(nil).Times(1)

	err := svc.SubmitReview(context.Background(), validReview)

	if err != nil {
		t.Errorf("Ekspektasi ulasan sukses disimpan, tapi mendapat error: %v", err)
	}
}
