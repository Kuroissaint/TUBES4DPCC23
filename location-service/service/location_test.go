package service_test

import (
	"testing"
	"location-service/service"
)

func TestCalculateDistance_Unit(t *testing.T) {
	lat1, lon1 := -6.200000, 106.816666
	lat2, lon2 := -6.210000, 106.826666

	t.Logf("Testing distance between: %f,%f and %f,%f", lat1, lon1, lat2, lon2)

	result := service.CalculateDistance(lat1, lon1, lat2, lon2)
	
	expected := 1.6 
	
	if result != expected {
		t.Errorf("Unit Test FAILED: Got %f, want %f", result, expected)
	}
}

func TestValidateCoordinates_Unit(t *testing.T) {
	invalidLat := 100.0
	invalidLon := 0.0
	
	isValid := service.ValidateCoordinates(invalidLat, invalidLon)

	if isValid {
		t.Errorf("Unit Test FAILED: Fungsi validasi harusnya me-return false untuk latitude 100")
	}
}
