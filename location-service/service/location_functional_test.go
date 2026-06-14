//go:build functional
// +build functional

package service_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq" 
)

func TestDatabaseConnection_Functional(t *testing.T) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "db_location" 
	}

	dbUser := "postgres"
	dbPass := "123"
	dbHost := "localhost" 

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
		dbHost, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal inisialisasi driver: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	
	if err != nil {
		t.Errorf("Functional Test FAILED: Tidak bisa terhubung ke database %s. Error: %v", dbName, err)
	}
}
