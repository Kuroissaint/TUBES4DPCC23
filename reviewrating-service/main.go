package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/handler"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/model"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/repository"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/service"
)

type dummyReviewRepo struct{}

func (r *dummyReviewRepo) SaveReview(ctx context.Context, rev model.Review) error {
	return nil
}

func main() {
	var repo repository.ReviewRepository = &dummyReviewRepo{}
	svc := service.NewReviewService(repo)
	hdl := handler.NewReviewHandler(svc)

	http.HandleFunc("/api/review/submit", hdl.SubmitReviewHandler)

	fmt.Println("ReviewRating Service running on :8083")
	http.ListenAndServe(":8083", nil)
}
