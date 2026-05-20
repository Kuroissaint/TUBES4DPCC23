package communication

import (
	"context"
)

// Struktur data pesan
type ChatMessage struct {
	OrderID       string
	SenderID      string
	MessageText   string
	AttachmentURL string
}

// Interface akses Database (Cassandra/MongoDB)
type ChatRepository interface {
	SaveMessage(ctx context.Context, msg ChatMessage) error
}