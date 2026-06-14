package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"location-service/model"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.WebResponse{
			Status:  "success",
			Message: "Location Service is running",
		})
	})

	router.POST("/track", func(c *gin.Context) {
		newID := uuid.New().String()
		response := model.WebResponse{
			Status: "success",
			Data: model.LocationData{
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