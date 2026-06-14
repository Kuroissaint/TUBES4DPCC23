//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"context"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/model"
)

type ReviewRepository interface {
	SaveReview(ctx context.Context, review model.Review) error
}
