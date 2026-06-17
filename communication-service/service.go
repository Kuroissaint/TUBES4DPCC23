package main

import (
	"context"
	"errors"
)

type ChatService struct {
	repo *ChatRepository
}

func NewChatService(repo *ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) SendMessage(ctx context.Context, msg Message) error {
	if msg.SenderID == "" || msg.ReceiverID == "" {
		return errors.New("pengirim dan penerima harus diisi")
	}
	if msg.Text == "" {
		return errors.New("pesan tidak boleh kosong")
	}
	
	return s.repo.SaveMessage(ctx, msg)
}