package main

import (
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

// CalculateDistance: Rumus Haversine buat dapet ~1.5km
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	
	// Hasil jarak antara -6.2, 106.816666 dan -6.21, 106.826666 adalah sekitar 1.5 km
	return math.Round(R*c*10) / 10 
}

// ValidateCoordinates: Filter koordinat biar gak ngawur
func ValidateCoordinates(lat, lon float64) bool {
	return lat >= -90 && lat <= 90 && lon >= -180 && lon <= 180
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, WebResponse{
			Status:  "success",
			Message: "Location Service is running",
		})
	})

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

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8002"
	}
	router.Run(":" + port)
}