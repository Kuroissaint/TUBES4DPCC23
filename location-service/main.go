package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"location-service/handler"
	"location-service/repository"
	"location-service/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 1. KONEKSI KE MONGODB-SERVICE KUBERNETES
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "mongodb://mongodb-service:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Gagal koneksi ke MongoDB: %v", err)
	}

	db := client.Database("tubes_location")

	// 2. INSIALISASI LAYER REPOSITORY, SERVICE, & HANDLER
	repo := repository.NewMongoLocationRepository(db)
	svc := service.NewLocationService(repo)
	hdl := handler.NewLocationHandler(svc)

	// 3. BALIKIN KE RUTE PANJANG ASLI BAWAAN (Biar lurus sama Gateway lu)
	http.HandleFunc("/location/distance", hdl.CalculateDistanceHandler)
    http.HandleFunc("/location/update", hdl.UpdateLocationHandler)
    http.HandleFunc("/location/nearby", hdl.GetNearbyDriversHandler)
	// 4. RUTE CATCH-ALL UNTUK CEK KESEHATAN SERVICE
	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"success","message":"Location Service is running smoothly with MongoDB"}`))
    })

	// 5. JALANKAN SERVER
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8002"
	}

	fmt.Printf("Location Service running on :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}