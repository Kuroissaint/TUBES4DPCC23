package main

import (
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
)

func TestFunctionalDBTransportConnection(t *testing.T) {
	// Menggunakan kredensial standar yang disepakati
	connStr := "user=postgres password=password dbname=db_transport_order sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal inisialisasi koneksi: %v", err)
	}
	defer db.Close()

	// Cek koneksi ping (Pasti akan failed jika DB belum jalan, sesuai ekspektasi dosen)
	err = db.Ping()
	if err != nil {
		t.Errorf("Functional Test Failed: Database db_transport_order belum siap atau belum di-setup: %v", err)
	}
}