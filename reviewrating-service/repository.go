package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Struktur data ulasan yang akan masuk ke JSON dan MongoDB (BSON)
type Review struct {
	OrderID  string `json:"order_id" bson:"order_id"`
	UserID   string `json:"user_id" bson:"user_id"`
	DriverID string `json:"driver_id" bson:"driver_id"`
	Rating   int    `json:"rating" bson:"rating"`
	Comment  string `json:"comment" bson:"comment"`
}

// Langsung implementasi fungsi database
type ReviewRepository struct {
	collection *mongo.Collection
}

func NewReviewRepository(db *mongo.Database) *ReviewRepository {
	return &ReviewRepository{
		collection: db.Collection("reviews"), // Nama tabel/koleksinya
	}
}

func (r *ReviewRepository) SaveReview(ctx context.Context, review Review) error {
	_, err := r.collection.InsertOne(ctx, review)
	return err
}