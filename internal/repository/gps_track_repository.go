package repository

import (
	"context"
	"time"

	"github.com/chats/sailing-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type gpsTrackRepository struct {
	collection *mongo.Collection
}

// NewGPSTrackRepository creates a new GPS track repository
func NewGPSTrackRepository(db *mongo.Database) domain.GPSTrackRepository {
	return &gpsTrackRepository{
		collection: db.Collection("gps_tracks"),
	}
}

func (r *gpsTrackRepository) CreateGPSTrack(ctx context.Context, track *domain.GPSTrack) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, track)
	if err != nil {
		return err
	}

	track.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *gpsTrackRepository) CreateGPSTracksBatch(ctx context.Context, tracks []*domain.GPSTrack) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Convert to []interface{} for InsertMany
	docs := make([]interface{}, len(tracks))
	for i, track := range tracks {
		docs[i] = track
	}

	results, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	// Update IDs in the original slice
	for i, id := range results.InsertedIDs {
		tracks[i].ID = id.(primitive.ObjectID)
	}

	return nil
}

func (r *gpsTrackRepository) GetGPSTracksByVoyageID(ctx context.Context, voyageID string) ([]*domain.GPSTrack, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"voyage_id": voyageID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tracks []*domain.GPSTrack
	if err = cursor.All(ctx, &tracks); err != nil {
		return nil, err
	}

	return tracks, nil
}
