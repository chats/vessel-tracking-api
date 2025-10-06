package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Voyage represents a sailing voyage
type Voyage struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	VoyageID      string             `json:"voyage_id" bson:"voyage_id"`
	ShipID        string             `json:"ship_id" bson:"ship_id"`
	ShipName      string             `json:"ship_name" bson:"ship_name"`
	DeparturePort string             `json:"departure_port" bson:"departure_port"`
	ArrivalPort   string             `json:"arrival_port,omitempty" bson:"arrival_port,omitempty"`
	DepartureTime time.Time          `json:"departure_time" bson:"departure_time"`
	ArrivalTime   *time.Time         `json:"arrival_time,omitempty" bson:"arrival_time,omitempty"`
	Status        string             `json:"status" bson:"status"` // "in_progress", "completed", "cancelled"
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}

// Checkpoint represents a checkpoint during a voyage
type Checkpoint struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	VoyageID    string             `json:"voyage_id" bson:"voyage_id"`
	Location    Location           `json:"location" bson:"location"`
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Weather     *WeatherInfo       `json:"weather,omitempty" bson:"weather,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// GPSTrack represents GPS tracking data
type GPSTrack struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	VoyageID  string             `json:"voyage_id" bson:"voyage_id"`
	Location  Location           `json:"location" bson:"location"`
	Speed     float64            `json:"speed" bson:"speed"`                           // knots
	Heading   float64            `json:"heading" bson:"heading"`                       // degrees
	Altitude  float64            `json:"altitude,omitempty" bson:"altitude,omitempty"` // meters
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// Location represents geographical coordinates
type Location struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

// WeatherInfo represents weather conditions at a checkpoint
type WeatherInfo struct {
	Temperature float64 `json:"temperature,omitempty" bson:"temperature,omitempty"` // celsius
	WindSpeed   float64 `json:"wind_speed,omitempty" bson:"wind_speed,omitempty"`   // knots
	WindDir     float64 `json:"wind_dir,omitempty" bson:"wind_dir,omitempty"`       // degrees
	WaveHeight  float64 `json:"wave_height,omitempty" bson:"wave_height,omitempty"` // meters
	Condition   string  `json:"condition,omitempty" bson:"condition,omitempty"`     // e.g., "clear", "cloudy", "rainy"
}

// VoyageWithDetails represents a voyage with its checkpoints and GPS tracks
type VoyageWithDetails struct {
	Voyage      *Voyage       `json:"voyage"`
	Checkpoints []*Checkpoint `json:"checkpoints"`
	GPSTracks   []*GPSTrack   `json:"gps_tracks"`
}
