package main

import (
	"context"
	"testing"
)

func TestSendMessage_Validation(t *testing.T) {
	service := NewChatService(nil)

	// Skenario 1: Teks pesan kosong
	invalidMsg := Message{
		SenderID:   "USR-1",
		ReceiverID: "USR-2",
		Text:       "", // Sengaja kosong
	}
	err := service.SendMessage(context.Background(), invalidMsg)
	if err == nil || err.Error() != "pesan tidak boleh kosong" {
		t.Errorf("Ekspektasi error pesan kosong, tapi mendapat: %v", err)
	}

	// Skenario 2: Sender/Receiver kosong
	invalidSender := Message{
		SenderID:   "", // Sengaja kosong
		ReceiverID: "USR-2",
		Text:       "Halo bro",
	}
	err2 := service.SendMessage(context.Background(), invalidSender)
	if err2 == nil || err2.Error() != "pengirim dan penerima harus diisi" {
		t.Errorf("Ekspektasi error pengirim kosong, tapi mendapat: %v", err2)
	}
}