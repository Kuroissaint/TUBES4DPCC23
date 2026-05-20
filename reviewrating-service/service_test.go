package reviewrating_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/mocks"
)

func TestSubmitReview_Failed_RatingTooHigh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockReviewRepository(ctrl)
	service := reviewrating.NewReviewService(mockRepo)

	invalidReview := reviewrating.Review{
		OrderID:  "ORD-111",
		UserID:   "USR-222",
		DriverID: "DRV-333",
		Rating:   6, // Sengaja diisi 6 agar memicu error
		Comment:  "Driver ngebut banget",
	}

	// Database tidak boleh dipanggil karena skenario gagal di validasi rating
	mockRepo.EXPECT().SaveReview(gomock.Any(), gomock.Any()).Times(0)

	err := service.SubmitReview(context.Background(), invalidReview)

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
	service := reviewrating.NewReviewService(mockRepo)

	validReview := reviewrating.Review{
		OrderID:  "ORD-111",
		UserID:   "USR-222",
		DriverID: "DRV-333",
		Rating:   5, // Valid
		Comment:  "Pelayanan sangat ramah dan memuaskan!",
	}

	// Database harus terpanggil 1 kali dan mengembalikan sukses (nil)
	mockRepo.EXPECT().SaveReview(gomock.Any(), validReview).Return(nil).Times(1)

	err := service.SubmitReview(context.Background(), validReview)

	if err != nil {
		t.Errorf("Ekspektasi ulasan sukses disimpan, tapi mendapat error: %v", err)
	}
}