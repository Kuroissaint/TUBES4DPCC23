package service

import (
	"errors"
	"dispatch-service/model"
)

// FindNearestDriver merepresentasikan logika pencarian driver terdekat
func FindNearestDriver(lat, lng float64) (model.Driver, error) {
	if lat == 0 && lng == 0 {
		return model.Driver{}, errors.New("lokasi tidak valid")
	}
	
	return model.Driver{
		ID:     "DRV-001",
		Name:   "Alex Marquez",
		Lat:    lat,
		Lng:    lng,
		Status: "available",
	}, nil
}

// AssignDriver mengubah status penugasan driver
func AssignDriver(driverID string, status string) (model.Driver, error) {
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
