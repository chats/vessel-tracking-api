package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/chats/sailing-backend/internal/domain"
)

// GPSTrackUseCase handles GPS track business logic
type GPSTrackUseCase struct {
	gpsTrackRepo domain.GPSTrackRepository
	voyageRepo   domain.VoyageRepository
}

// NewGPSTrackUseCase creates a new GPSTrackUseCase
func NewGPSTrackUseCase(gpsTrackRepo domain.GPSTrackRepository, voyageRepo domain.VoyageRepository) *GPSTrackUseCase {
	return &GPSTrackUseCase{
		gpsTrackRepo: gpsTrackRepo,
		voyageRepo:   voyageRepo,
	}
}

// CreateGPSTrack creates a new GPS track
func (uc *GPSTrackUseCase) CreateGPSTrack(ctx context.Context, track *domain.GPSTrack) error {
	if track.VoyageID == "" {
		return errors.New("voyage_id is required")
	}

	// Verify voyage exists
	_, err := uc.voyageRepo.GetVoyageByVoyageID(ctx, track.VoyageID)
	if err != nil {
		return errors.New("voyage not found")
	}

	track.CreatedAt = time.Now()
	if track.Timestamp.IsZero() {
		track.Timestamp = time.Now()
	}

	return uc.gpsTrackRepo.CreateGPSTrack(ctx, track)
}

// CreateGPSTracksBatch creates multiple GPS tracks
func (uc *GPSTrackUseCase) CreateGPSTracksBatch(ctx context.Context, tracks []*domain.GPSTrack) error {
	if len(tracks) == 0 {
		return errors.New("no GPS tracks provided")
	}

	// Validate and set timestamps
	for _, track := range tracks {
		if track.VoyageID == "" {
			return errors.New("voyage_id is required for all GPS tracks")
		}

		track.CreatedAt = time.Now()
		if track.Timestamp.IsZero() {
			track.Timestamp = time.Now()
		}
	}

	return uc.gpsTrackRepo.CreateGPSTracksBatch(ctx, tracks)
}
