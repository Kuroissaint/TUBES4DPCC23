package model

// Struktur GeoJSON untuk MongoDB
type GeoJson struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}

// Struktur Driver yang disimpan di DB
type DriverLocationDoc struct {
	DriverID  string  `bson:"driver_id" json:"driver_id"`
	Location  GeoJson `bson:"location" json:"location"`
	UpdatedAt string  `bson:"updated_at" json:"updated_at"`
}