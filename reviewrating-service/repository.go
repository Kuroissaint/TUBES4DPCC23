package reviewrating

import (
	"context"
)

// Struktur data ulasan
type Review struct {
	OrderID    string
	UserID     string
	DriverID   string
	Rating     int    // Harus bernilai 1 sampai 5
	Comment    string
}

// Interface yang akan di-mock oleh gomock
type ReviewRepository interface {
	SaveReview(ctx context.Context, review Review) error
}