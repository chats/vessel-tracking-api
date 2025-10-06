package handler

import (
	"github.com/chats/sailing-backend/internal/domain"
	"github.com/chats/sailing-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// CheckpointHandler handles checkpoint-related HTTP requests
type CheckpointHandler struct {
	checkpointUseCase *usecase.CheckpointUseCase
}

// NewCheckpointHandler creates a new checkpoint handler
func NewCheckpointHandler(checkpointUseCase *usecase.CheckpointUseCase) *CheckpointHandler {
	return &CheckpointHandler{
		checkpointUseCase: checkpointUseCase,
	}
}

// CreateCheckpoint creates a single checkpoint
func (h *CheckpointHandler) CreateCheckpoint(c *fiber.Ctx) error {
	var checkpoint domain.Checkpoint
	if err := c.BodyParser(&checkpoint); err != nil {
		log.Error().Err(err).Msg("Failed to parse checkpoint request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.checkpointUseCase.CreateCheckpoint(c.Context(), &checkpoint); err != nil {
		log.Error().Err(err).Msg("Failed to create checkpoint")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().
		Str("voyage_id", checkpoint.VoyageID).
		Str("checkpoint_id", checkpoint.ID.Hex()).
		Msg("Checkpoint created")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "checkpoint created successfully",
		"data":    checkpoint,
	})
}

// CreateCheckpointsBatch creates multiple checkpoints
func (h *CheckpointHandler) CreateCheckpointsBatch(c *fiber.Ctx) error {
	var checkpoints []*domain.Checkpoint
	if err := c.BodyParser(&checkpoints); err != nil {
		log.Error().Err(err).Msg("Failed to parse checkpoints batch request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.checkpointUseCase.CreateCheckpointsBatch(c.Context(), checkpoints); err != nil {
		log.Error().Err(err).Msg("Failed to create checkpoints batch")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().Int("count", len(checkpoints)).Msg("Checkpoints batch created")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "checkpoints created successfully",
		"data":    checkpoints,
		"count":   len(checkpoints),
	})
}
