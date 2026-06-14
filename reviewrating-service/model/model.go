package model

type Review struct {
	OrderID    string
	UserID     string
	DriverID   string
	Rating     int    // Harus bernilai 1 sampai 5
	Comment    string
}
