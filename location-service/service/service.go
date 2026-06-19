package service

import (
	"context"
	"math"
	"location-service/model"
	"location-service/repository"
)

type LocationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) *LocationService {
	return &LocationService{repo: repo}
}

func (s *LocationService) ValidateCoordinates(lat, lon float64) bool {
	return lat >= -90 && lat <= 90 && lon >= -180 && lon <= 180
}

func (s *LocationService) CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// Rumus Haversine asli bawaan kodemu tetap di sini
	const R = 6371000 // meter
	phi1 := lat1 * math.Pi / 180
	phi2 := lat2 * math.Pi / 180
	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// Fungsi Baru panggil Repo MongoDB
func (s *LocationService) SaveDriverLocation(ctx context.Context, driverID string, lat, lon float64) error {
	return s.repo.SaveLocation(ctx, driverID, lat, lon)
}

func (s *LocationService) FindNearbyDrivers(ctx context.Context, lat, lon, radius float64) ([]model.DriverLocationDoc, error) {
	return s.repo.GetNearby(ctx, lat, lon, radius)
}