package main

import (
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
)

func TestFunctionalDBShoppingConnection(t *testing.T) {
	connStr := "user=postgres password=password dbname=db_shopping_order sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal inisialisasi koneksi: %v", err)
	}
	defer db.Close()

	// Pasti akan failed di tahap ini
	err = db.Ping()
	if err != nil {
		t.Errorf("Functional Test Failed: Database db_shopping_order belum siap atau belum di-setup: %v", err)
	}
}