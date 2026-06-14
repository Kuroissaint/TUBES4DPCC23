//go:generate mockgen -source=repository.go -destination=../mocks/mock_repository.go -package=mocks
package repository

import (
	"context"
	"github.com/kuroissaint/tubes2dpcc/communication-service/model"
)

type ChatRepository interface {
	SaveMessage(ctx context.Context, msg model.ChatMessage) error
}
