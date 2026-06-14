package service

import (
	"context"
	"errors"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/model"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/repository"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) SubmitReview(ctx context.Context, rev model.Review) error {
	if rev.OrderID == "" || rev.UserID == "" || rev.DriverID == "" {
		return errors.New("order_id, user_id, dan driver_id wajib diisi")
	}

	if rev.Rating < 1 || rev.Rating > 5 {
		return errors.New("rating tidak valid: harus berada di antara rentang angka 1 sampai 5")
	}

	err := s.repo.SaveReview(ctx, rev)
	if err != nil {
		return err
	}

	return nil
}
