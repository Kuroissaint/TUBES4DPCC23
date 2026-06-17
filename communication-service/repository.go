package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	SenderID   string    `json:"sender_id" bson:"sender_id"`
	ReceiverID string    `json:"receiver_id" bson:"receiver_id"`
	Text       string    `json:"text" bson:"text"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

type ChatRepository struct {
	collection *mongo.Collection
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	return &ChatRepository{
		collection: db.Collection("messages"),
	}
}

func (r *ChatRepository) SaveMessage(ctx context.Context, msg Message) error {
	msg.CreatedAt = time.Now() // Otomatis catat waktu pengiriman
	_, err := r.collection.InsertOne(ctx, msg)
	return err
}