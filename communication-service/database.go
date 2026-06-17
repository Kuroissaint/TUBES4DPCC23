package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Langsung pakai localhost untuk tes lokal
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@mongodb-cluster:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Gagal konek ke MongoDB:", err)
	}

	// Gunakan database bernama "db_chat"
	return client.Database("db_chat")
}