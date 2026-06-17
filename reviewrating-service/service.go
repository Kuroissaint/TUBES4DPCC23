package main

import (
	"context"
	"errors"
)

type ReviewService struct {
	repo *ReviewRepository // Terhubung langsung tanpa interface
}

func NewReviewService(repo *ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) SubmitReview(ctx context.Context, rev Review) error {
	if rev.OrderID == "" || rev.UserID == "" || rev.DriverID == "" {
		return errors.New("order_id, user_id, dan driver_id wajib diisi")
	}
	if rev.Rating < 1 || rev.Rating > 5 {
		return errors.New("rating tidak valid: harus berada di antara 1 sampai 5")
	}
	
	return s.repo.SaveReview(ctx, rev)
}