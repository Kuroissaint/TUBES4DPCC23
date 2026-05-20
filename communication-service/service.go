package communication

import (
	"context"
	"errors"
)

type ChatService struct {
	repo ChatRepository
}

func NewChatService(repo ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

// Fungsi memproses chat masuk
func (s *ChatService) ProcessIncomingMessage(ctx context.Context, msg ChatMessage) error {
	// 1. Validasi
	if msg.OrderID == "" || msg.SenderID == "" {
		return errors.New("order_id dan sender_id tidak boleh kosong")
	}
	if msg.MessageText == "" && msg.AttachmentURL == "" {
		return errors.New("pesan tidak valid: harus ada teks atau gambar")
	}

	// 2. Simpan ke database
	err := s.repo.SaveMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}