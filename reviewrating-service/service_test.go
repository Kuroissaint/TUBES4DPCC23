package main

import (
	"context"
	"testing"
)

// Tes ini hanya menguji logika validasi dasar tanpa perlu konek ke database atau layanan lain
func TestSubmitReview_Validation(t *testing.T) {
	// Buat service dengan repo kosong (nil)
	service := NewReviewService(nil)

	// Skenario 1: Sengaja mengosongkan OrderID agar gagal
	invalidReview := Review{
		OrderID:  "", // Sengaja kosong
		UserID:   "USR-222",
		DriverID: "DRV-333",
		Rating:   5,
		Comment:  "Bagus",
	}

	err := service.SubmitReview(context.Background(), invalidReview)
	if err == nil {
		t.Errorf("Ekspektasi error karena order_id kosong, tapi malah sukses")
	}

	// Skenario 2: Rating di luar batas (misal 6)
	invalidRatingReview := Review{
		OrderID:  "ORD-111",
		UserID:   "USR-222",
		DriverID: "DRV-333",
		Rating:   6, // Tidak valid
		Comment:  "Terlalu bagus",
	}

	err2 := service.SubmitReview(context.Background(), invalidRatingReview)
	if err2 == nil {
		t.Errorf("Ekspektasi error karena rating 6, tapi malah sukses")
	}
}