package repository

import (
	"context"
	"time"

	"github.com/chats/sailing-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type checkpointRepository struct {
	collection *mongo.Collection
}

// NewCheckpointRepository creates a new checkpoint repository
func NewCheckpointRepository(db *mongo.Database) domain.CheckpointRepository {
	return &checkpointRepository{
		collection: db.Collection("checkpoints"),
	}
}

func (r *checkpointRepository) CreateCheckpoint(ctx context.Context, checkpoint *domain.Checkpoint) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, checkpoint)
	if err != nil {
		return err
	}

	checkpoint.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *checkpointRepository) CreateCheckpointsBatch(ctx context.Context, checkpoints []*domain.Checkpoint) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Convert to []interface{} for InsertMany
	docs := make([]interface{}, len(checkpoints))
	for i, checkpoint := range checkpoints {
		docs[i] = checkpoint
	}

	results, err := r.collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	// Update IDs in the original slice
	for i, id := range results.InsertedIDs {
		checkpoints[i].ID = id.(primitive.ObjectID)
	}

	return nil
}

func (r *checkpointRepository) GetCheckpointsByVoyageID(ctx context.Context, voyageID string) ([]*domain.Checkpoint, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"voyage_id": voyageID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var checkpoints []*domain.Checkpoint
	if err = cursor.All(ctx, &checkpoints); err != nil {
		return nil, err
	}

	return checkpoints, nil
}
