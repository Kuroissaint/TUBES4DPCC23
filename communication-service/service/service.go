package service

import (
	"context"
	"errors"
	"github.com/kuroissaint/tubes2dpcc/communication-service/model"
	"github.com/kuroissaint/tubes2dpcc/communication-service/repository"
)

type ChatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) ProcessIncomingMessage(ctx context.Context, msg model.ChatMessage) error {
	if msg.OrderID == "" || msg.SenderID == "" {
		return errors.New("order_id dan sender_id tidak boleh kosong")
	}
	if msg.MessageText == "" && msg.AttachmentURL == "" {
		return errors.New("pesan tidak valid: harus ada teks atau gambar")
	}

	err := s.repo.SaveMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}
