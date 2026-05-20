package main

import (
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Struktur respon standar
type WebResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type LocationData struct {
	UserID    string  `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// CalculateDistance menghitung jarak antar dua titik (Haversine Formula)
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius bumi dalam km
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// ValidateCoordinates memastikan latitude/longitude dalam range valid
func ValidateCoordinates(lat, lon float64) bool {
	if lat < -90 || lat > 90 {
		return false
	}
	if lon < -180 || lon > 180 {
		return false
	}
	return true
}

func main() {
	router := gin.Default()

	// 1. Endpoint Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, WebResponse{
			Status:  "success",
			Message: "Location Service is running",
		})
	})

	// 2. Endpoint Tracking
	router.POST("/track", func(c *gin.Context) {
		newID := uuid.New().String()
		response := WebResponse{
			Status: "success",
			Data: LocationData{
				UserID:    newID,
				Latitude:  -6.200000,
				Longitude: 106.816666,
			},
			Message: "Location tracked successfully",
		}
		c.JSON(http.StatusOK, response)
	})

	// Mengambil port dari env atau default 8002
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8002"
	}

	fmt.Println("Location Service running on port:", port)
	router.Run(":" + port)
}