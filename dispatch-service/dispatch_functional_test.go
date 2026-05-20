//go:build functional
// +build functional

package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq" // Driver postgres
)

// Functional Test: Menguji koneksi ke database db_dispatch
func TestDispatchDBConnection_Functional(t *testing.T) {
	// Mengambil info DB dari environment
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "db_dispatch" // Nama DB khusus service ini
	}

	dbUser := "PostgresLocal"
	dbPass := "123"
	dbHost := "host.docker.internal"

	// Membuat string koneksi
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
		dbHost, dbUser, dbPass, dbName)

	// Mencoba membuka koneksi
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal inisialisasi driver database: %v", err)
	}
	defer db.Close()

	// Melakukan ping ke database
	err = db.Ping()
	
	// Sengaja dibuat FAILED karena database belum ada di tahap dua ini
	if err != nil {
		t.Errorf("Functional Test FAILED: Tidak bisa terhubung ke database %s. Error: %v", dbName, err)
	}
}