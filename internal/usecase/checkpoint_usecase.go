package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/chats/sailing-backend/internal/domain"
)

// CheckpointUseCase handles checkpoint business logic
type CheckpointUseCase struct {
	checkpointRepo domain.CheckpointRepository
	voyageRepo     domain.VoyageRepository
}

// NewCheckpointUseCase creates a new CheckpointUseCase
func NewCheckpointUseCase(checkpointRepo domain.CheckpointRepository, voyageRepo domain.VoyageRepository) *CheckpointUseCase {
	return &CheckpointUseCase{
		checkpointRepo: checkpointRepo,
		voyageRepo:     voyageRepo,
	}
}

// CreateCheckpoint creates a new checkpoint
func (uc *CheckpointUseCase) CreateCheckpoint(ctx context.Context, checkpoint *domain.Checkpoint) error {
	if checkpoint.VoyageID == "" {
		return errors.New("voyage_id is required")
	}

	// Verify voyage exists
	_, err := uc.voyageRepo.GetVoyageByVoyageID(ctx, checkpoint.VoyageID)
	if err != nil {
		return errors.New("voyage not found")
	}

	checkpoint.CreatedAt = time.Now()
	if checkpoint.Timestamp.IsZero() {
		checkpoint.Timestamp = time.Now()
	}

	return uc.checkpointRepo.CreateCheckpoint(ctx, checkpoint)
}

// CreateCheckpointsBatch creates multiple checkpoints
func (uc *CheckpointUseCase) CreateCheckpointsBatch(ctx context.Context, checkpoints []*domain.Checkpoint) error {
	if len(checkpoints) == 0 {
		return errors.New("no checkpoints provided")
	}

	// Validate and set timestamps
	for _, checkpoint := range checkpoints {
		if checkpoint.VoyageID == "" {
			return errors.New("voyage_id is required for all checkpoints")
		}

		checkpoint.CreatedAt = time.Now()
		if checkpoint.Timestamp.IsZero() {
			checkpoint.Timestamp = time.Now()
		}
	}

	return uc.checkpointRepo.CreateCheckpointsBatch(ctx, checkpoints)
}
