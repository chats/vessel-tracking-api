package domain

import (
	"context"
)

// VoyageRepository defines the interface for voyage data operations
type VoyageRepository interface {
	CreateVoyage(ctx context.Context, voyage *Voyage) error
	UpdateVoyage(ctx context.Context, voyage *Voyage) error
	GetVoyageByID(ctx context.Context, id string) (*Voyage, error)
	GetAllVoyages(ctx context.Context, limit, offset int) ([]*Voyage, error)
	GetVoyageByVoyageID(ctx context.Context, voyageID string) (*Voyage, error)
}

// CheckpointRepository defines the interface for checkpoint data operations
type CheckpointRepository interface {
	CreateCheckpoint(ctx context.Context, checkpoint *Checkpoint) error
	CreateCheckpointsBatch(ctx context.Context, checkpoints []*Checkpoint) error
	GetCheckpointsByVoyageID(ctx context.Context, voyageID string) ([]*Checkpoint, error)
}

// GPSTrackRepository defines the interface for GPS track data operations
type GPSTrackRepository interface {
	CreateGPSTrack(ctx context.Context, track *GPSTrack) error
	CreateGPSTracksBatch(ctx context.Context, tracks []*GPSTrack) error
	GetGPSTracksByVoyageID(ctx context.Context, voyageID string) ([]*GPSTrack, error)
}
