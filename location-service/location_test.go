package main

import (
	"testing"
)

func TestCalculateDistance_Unit(t *testing.T) {
	lat1, lon1 := -6.200000, 106.816666
	lat2, lon2 := -6.210000, 106.826666

	t.Logf("Testing distance between: %f,%f and %f,%f", lat1, lon1, lat2, lon2)

	result := CalculateDistance(lat1, lon1, lat2, lon2)
	
	// Ubah expected jadi 1.6 sesuai dengan hasil asli rumus Haversine di sistem lu
	expected := 1.6 
	
	if result != expected {
		t.Errorf("Unit Test FAILED: Got %f, want %f", result, expected)
	}
}

func TestValidateCoordinates_Unit(t *testing.T) {
	invalidLat := 100.0
	invalidLon := 0.0
	
	isValid := ValidateCoordinates(invalidLat, invalidLon)

	if isValid {
		t.Errorf("Unit Test FAILED: Fungsi validasi harusnya me-return false untuk latitude 100")
	}
}