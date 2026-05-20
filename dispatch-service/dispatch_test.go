package main

import (
	"testing"
)

func TestFindNearestDriver_Unit(t *testing.T) {
	lat := -6.200000
	lng := 106.816666
	t.Logf("Mencari driver terdekat untuk lokasi: %f, %f", lat, lng)

	// Memanggil fungsi asli dari main.go
	driver, err := FindNearestDriver(lat, lng)
	if err != nil {
		t.Fatalf("Fungsi error: %v", err)
	}

	// Memastikan data driver terdekat berhasil ditemukan
	if driver.ID == "" {
		t.Errorf("Unit Test FAILED: ID Driver kosong")
	}
	if driver.Name != "Alex Marquez" {
		t.Errorf("Unit Test FAILED: Nama driver tidak sesuai, dapat: %s", driver.Name)
	}
}

func TestAssignDriverStatus_Unit(t *testing.T) {
	driverID := "DRV-001"
	status := "pending"

	// Memanggil fungsi asli dari main.go
	driver, err := AssignDriver(driverID, status)
	if err != nil {
		t.Fatalf("Fungsi error: %v", err)
	}

	// Memastikan status berhasil berubah menjadi 'assigned' dan tidak tertahan di 'pending'
	if driver.Status == "pending" {
		t.Fatalf("Unit Test FAILED: Logika penugasan gagal, status tetap: pending")
	}

	if driver.Status != "assigned" {
		t.Errorf("Expect status 'assigned', got '%s'", driver.Status)
	}
}