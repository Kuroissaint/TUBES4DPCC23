package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"dispatch-service/model"
	"dispatch-service/repository"
)

type DispatchService struct {
	repo repository.DispatchRepository
}

func NewDispatchService(repo repository.DispatchRepository) *DispatchService {
	return &DispatchService{repo: repo}
}

// FindNearestDriver merepresentasikan logika pencarian driver terdekat menggunakan Location Service
func (s *DispatchService) FindNearestDriver(lat, lng float64) (model.Driver, error) {
	if lat == 0 && lng == 0 {
		return model.Driver{}, errors.New("lokasi tidak valid")
	}

	locationURL := fmt.Sprintf("http://location-service-service:8002/location/nearby?lat=%f&lon=%f", lat, lng)
	resp, err := http.Get(locationURL)
	if err != nil {
		fmt.Println("[WARNING] Gagal menghubungi Location Service:", err)
		return model.Driver{
			ID:     "DRV-001",
			Name:   "Alex Marquez (Fallback)",
			Lat:    lat,
			Lng:    lng,
			Status: "available",
		}, nil
	}
	defer resp.Body.Close()

	var locResp struct {
		Status  string `json:"status"`
		Drivers []struct {
			DriverID string  `json:"driver_id"`
			Latitude float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"drivers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&locResp); err == nil && len(locResp.Drivers) > 0 {
		nearest := locResp.Drivers[0]
		return model.Driver{
			ID:     nearest.DriverID,
			Name:   "Dynamic Driver " + nearest.DriverID,
			Lat:    nearest.Latitude,
			Lng:    nearest.Longitude,
			Status: "available",
		}, nil
	}

	return model.Driver{
		ID:     "DRV-001",
		Name:   "Alex Marquez (Fallback)",
		Lat:    lat,
		Lng:    lng,
		Status: "available",
	}, nil
}

// AssignDriver mengubah status penugasan driver
func (s *DispatchService) AssignDriver(driverID string, status string) (model.Driver, error) {
	finalStatus := status
	if status == "pending" || status == "" {
		finalStatus = "assigned"
	}

	return model.Driver{
		ID:     driverID,
		Name:   "Marc Marquez",
		Status: finalStatus,
	}, nil
}
