package main

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type ReviewService struct {
	repo *ReviewRepository // Terhubung langsung tanpa interface
}

func NewReviewService(repo *ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) SubmitReview(ctx context.Context, rev Review) error {
	// 1. Validasi input dasar
	if rev.OrderID == "" || rev.UserID == "" || rev.DriverID == "" {
		return errors.New("order_id, user_id, dan driver_id wajib diisi")
	}
	if rev.Rating < 1 || rev.Rating > 5 {
		return errors.New("rating tidak valid: harus berada di antara 1 sampai 5")
	}

	// === 2. TAMBAHAN KODE KOMUNIKASI ANTAR LAYANAN ===
	// Kita buat HTTP client dengan batas waktu (timeout) 5 detik. 
	// Ini penting agar layananmu tidak "hang" kalau layanan sebelah kebetulan lagi mati.
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Tembak ke URL service Kubernetes milik Shop Order.
	// (Ganti port 8080 dan jalur "/orders/" sesuai dengan kesepakatan tim)
	url := "http://shop-order-service:8080/orders/" + rev.OrderID

	resp, err := client.Get(url)
	if err != nil {
		// Error ini memblokir jika layanan order sama sekali tidak bisa dihubungi
		return errors.New("gagal memvalidasi pesanan ke layanan order: " + err.Error())
	}
	defer resp.Body.Close() // Wajib ditutup untuk mencegah kebocoran memori

	// Cek apakah balasan dari Shop Order sukses (Status 200 OK)
	// Jika bukan 200 (misal 404 Not Found), berarti pesanan fiktif atau belum selesai
	if resp.StatusCode != http.StatusOK {
		return errors.New("ditolak: pesanan tidak ditemukan atau status pesanan belum selesai")
	}
	// ================================================

	// 3. Jika semua validasi lulus, simpan ulasan ke database
	return s.repo.SaveReview(ctx, rev)
}