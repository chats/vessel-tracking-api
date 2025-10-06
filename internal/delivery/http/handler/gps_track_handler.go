package handler

import (
	"github.com/chats/sailing-backend/internal/domain"
	"github.com/chats/sailing-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// GPSTrackHandler handles GPS track-related HTTP requests
type GPSTrackHandler struct {
	gpsTrackUseCase *usecase.GPSTrackUseCase
}

// NewGPSTrackHandler creates a new GPS track handler
func NewGPSTrackHandler(gpsTrackUseCase *usecase.GPSTrackUseCase) *GPSTrackHandler {
	return &GPSTrackHandler{
		gpsTrackUseCase: gpsTrackUseCase,
	}
}

// CreateGPSTrack creates a single GPS track
func (h *GPSTrackHandler) CreateGPSTrack(c *fiber.Ctx) error {
	var track domain.GPSTrack
	if err := c.BodyParser(&track); err != nil {
		log.Error().Err(err).Msg("Failed to parse GPS track request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.gpsTrackUseCase.CreateGPSTrack(c.Context(), &track); err != nil {
		log.Error().Err(err).Msg("Failed to create GPS track")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().
		Str("voyage_id", track.VoyageID).
		Str("track_id", track.ID.Hex()).
		Msg("GPS track created")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "GPS track created successfully",
		"data":    track,
	})
}

// CreateGPSTracksBatch creates multiple GPS tracks
func (h *GPSTrackHandler) CreateGPSTracksBatch(c *fiber.Ctx) error {
	var tracks []*domain.GPSTrack
	if err := c.BodyParser(&tracks); err != nil {
		log.Error().Err(err).Msg("Failed to parse GPS tracks batch request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.gpsTrackUseCase.CreateGPSTracksBatch(c.Context(), tracks); err != nil {
		log.Error().Err(err).Msg("Failed to create GPS tracks batch")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().Int("count", len(tracks)).Msg("GPS tracks batch created")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "GPS tracks created successfully",
		"data":    tracks,
		"count":   len(tracks),
	})
}
