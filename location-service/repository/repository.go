package repository

import (
	"context"
	"fmt" // <-- Tambahkan ini untuk log
	"time"
	"location-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocationRepository interface {
	SaveLocation(ctx context.Context, driverID string, lat, lon float64) error
	GetNearby(ctx context.Context, lat, lon float64, radiusMeter float64) ([]model.DriverLocationDoc, error)
}

type MongoLocationRepository struct {
	collection *mongo.Collection
}

func NewMongoLocationRepository(db *mongo.Database) *MongoLocationRepository {
	coll := db.Collection("driver_locations")

	// === DI SINI TEMPATNYA, BANG! (AUTO-INDEX GEOSPATIAL) ===
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		indexModel := mongo.IndexModel{
			Keys: bson.M{"location": "2dsphere"}, // Mengunci field location agar bisa query $near
		}

		_, err := coll.Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			fmt.Printf("❌ Gagal membuat geospatial index: %v\n", err)
		} else {
			fmt.Println("✅ Geospatial index '2dsphere' berhasil diverifikasi/dibuat di MongoDB.")
		}
	}()
	// ========================================================

	return &MongoLocationRepository{
		collection: coll,
	}
}

// Menyimpan atau memperbarui lokasi driver (Upsert)
func (r *MongoLocationRepository) SaveLocation(ctx context.Context, driverID string, lat, lon float64) error {
	filter := bson.M{"driver_id": driverID}
	update := bson.M{
		"$set": bson.M{
			"location": bson.M{
				"type":        "Point",
				"coordinates": []float64{lon, lat}, // MongoDB menerima [longitude, latitude]
			},
			"updated_at": time.Now().Format(time.RFC3339),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

// Mencari driver terdekat dengan fitur Geospatial MongoDB ($near)
func (r *MongoLocationRepository) GetNearby(ctx context.Context, lat, lon float64, radiusMeter float64) ([]model.DriverLocationDoc, error) {
	query := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lon, lat},
				},
				"$maxDistance": radiusMeter,
			},
		},
	}

	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var drivers []model.DriverLocationDoc
	if err = cursor.All(ctx, &drivers); err != nil {
		return nil, err
	}

	return drivers, nil
}