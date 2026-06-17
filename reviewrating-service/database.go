package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
	// Beri batas waktu 10 detik untuk mencoba konek
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("==== HALO! SAYA SEDANG MENCOBA KONEK KE LOCALHOST ====")

	// Menyambung ke MongoDB di Kubernetes dengan username & password
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@mongodb-cluster:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Gagal konek ke MongoDB:", err)
	}

	// Otomatis membuat/menggunakan database bernama "db_review"
	return client.Database("db_review")
}