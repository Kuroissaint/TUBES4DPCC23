package reviewrating

import (
	"context"
	"errors"
)

type ReviewService struct {
	repo ReviewRepository
}

func NewReviewService(repo ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

// Fungsi untuk memproses ulasan baru masuk
func (s *ReviewService) SubmitReview(ctx context.Context, rev Review) error {
	// 1. Validasi Kelengkapan ID
	if rev.OrderID == "" || rev.UserID == "" || rev.DriverID == "" {
		return errors.New("order_id, user_id, dan driver_id wajib diisi")
	}

	// 2. Validasi Batas Rating (Target Utama Unit Test Tahap 2)
	if rev.Rating < 1 || rev.Rating > 5 {
		return errors.New("rating tidak valid: harus berada di antara rentang angka 1 sampai 5")
	}

	// 3. Simpan ke Database lewat Repository jika lolos validasi
	err := s.repo.SaveReview(ctx, rev)
	if err != nil {
		return err
	}

	return nil
}