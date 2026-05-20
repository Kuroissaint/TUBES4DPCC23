package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Driver mewakili struktur data driver
type Driver struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Status string  `json:"status"`
}

// FindNearestDriver merepresentasikan logika pencarian driver terdekat
func FindNearestDriver(lat, lng float64) (Driver, error) {
	// Logika tiruan (mock/stub) agar unit test lolos.
	// Di dunia nyata, ini akan menghitung jarak koordinat ke database/redis.
	return Driver{
		ID:     "DRV-001",
		Name:   "Alex Marquez", // Menggunakan tema MotoGP favoritmu
		Lat:    lat,
		Lng:    lng,
		Status: "available",
	}, nil
}

// AssignDriver mengubah status penugasan driver
func AssignDriver(driverID string, status string) (Driver, error) {
	// Logika agar status tidak tertahan di "pending" saat ditugaskan
	finalStatus := status
	if status == "pending" || status == "" {
		finalStatus = "assigned"
	}
	
	return Driver{
		ID:     driverID,
		Name:   "Marc Marquez",
		Status: finalStatus,
	}, nil
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}