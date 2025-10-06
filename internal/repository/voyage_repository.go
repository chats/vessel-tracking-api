package repository

import (
	"context"
	"errors"
	"time"

	"github.com/chats/sailing-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type voyageRepository struct {
	collection *mongo.Collection
}

// NewVoyageRepository creates a new voyage repository
func NewVoyageRepository(db *mongo.Database) domain.VoyageRepository {
	return &voyageRepository{
		collection: db.Collection("voyages"),
	}
}

func (r *voyageRepository) CreateVoyage(ctx context.Context, voyage *domain.Voyage) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, voyage)
	if err != nil {
		return err
	}

	voyage.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *voyageRepository) UpdateVoyage(ctx context.Context, voyage *domain.Voyage) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": voyage.ID}
	update := bson.M{
		"$set": bson.M{
			"arrival_port": voyage.ArrivalPort,
			"arrival_time": voyage.ArrivalTime,
			"status":       voyage.Status,
			"updated_at":   voyage.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("voyage not found")
	}

	return nil
}

func (r *voyageRepository) GetVoyageByID(ctx context.Context, id string) (*domain.Voyage, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid voyage ID")
	}

	var voyage domain.Voyage
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&voyage)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("voyage not found")
		}
		return nil, err
	}

	return &voyage, nil
}

func (r *voyageRepository) GetAllVoyages(ctx context.Context, limit, offset int) ([]*domain.Voyage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var voyages []*domain.Voyage
	if err = cursor.All(ctx, &voyages); err != nil {
		return nil, err
	}

	return voyages, nil
}

func (r *voyageRepository) GetVoyageByVoyageID(ctx context.Context, voyageID string) (*domain.Voyage, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var voyage domain.Voyage
	err := r.collection.FindOne(ctx, bson.M{"voyage_id": voyageID}).Decode(&voyage)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("voyage not found")
		}
		return nil, err
	}

	return &voyage, nil
}
