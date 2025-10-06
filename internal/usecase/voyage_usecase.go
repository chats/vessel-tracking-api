package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/chats/sailing-backend/internal/domain"
	"github.com/google/uuid"
)

// VoyageUseCase handles voyage business logic
type VoyageUseCase struct {
	voyageRepo     domain.VoyageRepository
	checkpointRepo domain.CheckpointRepository
	gpsTrackRepo   domain.GPSTrackRepository
}

// NewVoyageUseCase creates a new VoyageUseCase
func NewVoyageUseCase(voyageRepo domain.VoyageRepository, checkpointRepo domain.CheckpointRepository, gpsTrackRepo domain.GPSTrackRepository) *VoyageUseCase {
	return &VoyageUseCase{
		voyageRepo:     voyageRepo,
		checkpointRepo: checkpointRepo,
		gpsTrackRepo:   gpsTrackRepo,
	}
}

// DepartVoyage creates a new voyage with departure information
func (uc *VoyageUseCase) DepartVoyage(ctx context.Context, voyage *domain.Voyage) error {
	if voyage.ShipID == "" {
		return errors.New("ship_id is required")
	}
	if voyage.ShipName == "" {
		return errors.New("ship_name is required")
	}
	if voyage.DeparturePort == "" {
		return errors.New("departure_port is required")
	}

	// Generate voyage ID if not provided
	if voyage.VoyageID == "" {
		voyage.VoyageID = uuid.New().String()
	}

	voyage.Status = "in_progress"
	voyage.DepartureTime = time.Now()
	voyage.CreatedAt = time.Now()
	voyage.UpdatedAt = time.Now()

	return uc.voyageRepo.CreateVoyage(ctx, voyage)
}

// ArriveVoyage updates a voyage with arrival information
func (uc *VoyageUseCase) ArriveVoyage(ctx context.Context, voyageID, arrivalPort string) error {
	voyage, err := uc.voyageRepo.GetVoyageByVoyageID(ctx, voyageID)
	if err != nil {
		return err
	}

	if voyage.Status != "in_progress" {
		return errors.New("voyage is not in progress")
	}

	now := time.Now()
	voyage.ArrivalPort = arrivalPort
	voyage.ArrivalTime = &now
	voyage.Status = "completed"
	voyage.UpdatedAt = time.Now()

	return uc.voyageRepo.UpdateVoyage(ctx, voyage)
}

// GetAllVoyages retrieves all voyages with their checkpoints and GPS tracks
func (uc *VoyageUseCase) GetAllVoyages(ctx context.Context, limit, offset int) ([]*domain.VoyageWithDetails, error) {
	if limit <= 0 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	voyages, err := uc.voyageRepo.GetAllVoyages(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Fetch checkpoints and GPS tracks for each voyage
	result := make([]*domain.VoyageWithDetails, 0, len(voyages))
	for _, voyage := range voyages {
		checkpoints, err := uc.checkpointRepo.GetCheckpointsByVoyageID(ctx, voyage.VoyageID)
		if err != nil {
			// Log error but continue with empty checkpoints
			checkpoints = []*domain.Checkpoint{}
		}

		gpsTracks, err := uc.gpsTrackRepo.GetGPSTracksByVoyageID(ctx, voyage.VoyageID)
		if err != nil {
			// Log error but continue with empty GPS tracks
			gpsTracks = []*domain.GPSTrack{}
		}

		result = append(result, &domain.VoyageWithDetails{
			Voyage:      voyage,
			Checkpoints: checkpoints,
			GPSTracks:   gpsTracks,
		})
	}

	return result, nil
}

// GetVoyageByID retrieves a voyage by ID
func (uc *VoyageUseCase) GetVoyageByID(ctx context.Context, id string) (*domain.Voyage, error) {
	return uc.voyageRepo.GetVoyageByID(ctx, id)
}
