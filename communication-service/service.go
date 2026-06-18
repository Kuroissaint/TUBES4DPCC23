package main

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type ChatService struct {
	repo *ChatRepository
}

func NewChatService(repo *ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) SendMessage(ctx context.Context, msg Message) error {
	// 1. Validasi input dasar
	if msg.SenderID == "" || msg.ReceiverID == "" {
		return errors.New("pengirim dan penerima harus diisi")
	}
	if msg.Text == "" {
		return errors.New("pesan tidak boleh kosong")
	}

	// === 2. TAMBAHAN KODE KOMUNIKASI ANTAR LAYANAN ===
	// Membuat HTTP client dengan batas waktu (timeout) 5 detik demi keandalan sistem
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Langkah A: Validasi SenderID (Pengirim) ke account-service
	// (Silakan ganti port dan jalur endpoint-nya sesuai dengan setup milik teman kelompokmu)
	senderUrl := "http://account-service:8081/users/" + msg.SenderID
	respSender, err := client.Get(senderUrl)
	if err != nil {
		return errors.New("gagal memvalidasi pengirim ke layanan akun: " + err.Error())
	}
	defer respSender.Body.Close()

	if respSender.StatusCode != http.StatusOK {
		return errors.New("ditolak: pengirim (sender_id) tidak valid atau tidak ditemukan")
	}

	// Langkah B: Validasi ReceiverID (Penerima) ke account-service
	receiverUrl := "http://account-service:8081/users/" + msg.ReceiverID
	respReceiver, err := client.Get(receiverUrl)
	if err != nil {
		return errors.New("gagal memvalidasi penerima ke layanan akun: " + err.Error())
	}
	defer respReceiver.Body.Close()

	if respReceiver.StatusCode != http.StatusOK {
		return errors.New("ditolak: penerima (receiver_id) tidak valid atau tidak ditemukan")
	}
	// ================================================

	// 3. Jika pengirim dan penerima terbukti valid, simpan pesan ke MongoDB
	return s.repo.SaveMessage(ctx, msg)
}