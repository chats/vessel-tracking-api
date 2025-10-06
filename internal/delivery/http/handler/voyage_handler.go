package handler

import (
	"strconv"

	"github.com/chats/sailing-backend/internal/domain"
	"github.com/chats/sailing-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// VoyageHandler handles voyage-related HTTP requests
type VoyageHandler struct {
	voyageUseCase *usecase.VoyageUseCase
}

// NewVoyageHandler creates a new voyage handler
func NewVoyageHandler(voyageUseCase *usecase.VoyageUseCase) *VoyageHandler {
	return &VoyageHandler{
		voyageUseCase: voyageUseCase,
	}
}

// DepartRequest represents the depart request body
type DepartRequest struct {
	ShipID        string `json:"ship_id"`
	ShipName      string `json:"ship_name"`
	DeparturePort string `json:"departure_port"`
	VoyageID      string `json:"voyage_id,omitempty"`
}

// ArriveRequest represents the arrive request body
type ArriveRequest struct {
	VoyageID    string `json:"voyage_id"`
	ArrivalPort string `json:"arrival_port"`
}

// Depart handles voyage departure
func (h *VoyageHandler) Depart(c *fiber.Ctx) error {
	var req DepartRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("Failed to parse depart request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	voyage := &domain.Voyage{
		VoyageID:      req.VoyageID,
		ShipID:        req.ShipID,
		ShipName:      req.ShipName,
		DeparturePort: req.DeparturePort,
	}

	if err := h.voyageUseCase.DepartVoyage(c.Context(), voyage); err != nil {
		log.Error().Err(err).Msg("Failed to create voyage")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().Str("voyage_id", voyage.VoyageID).Msg("Voyage departed")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "voyage departed successfully",
		"data":    voyage,
	})
}

// Arrive handles voyage arrival
func (h *VoyageHandler) Arrive(c *fiber.Ctx) error {
	var req ArriveRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("Failed to parse arrive request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.voyageUseCase.ArriveVoyage(c.Context(), req.VoyageID, req.ArrivalPort); err != nil {
		log.Error().Err(err).Msg("Failed to update voyage arrival")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Info().Str("voyage_id", req.VoyageID).Msg("Voyage arrived")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "voyage arrived successfully",
	})
}

// GetAllVoyages retrieves all voyages
func (h *VoyageHandler) GetAllVoyages(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "100"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	voyages, err := h.voyageUseCase.GetAllVoyages(c.Context(), limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get voyages")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve voyages",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  voyages,
		"count": len(voyages),
	})
}

// GetVoyageByID retrieves a voyage by ID
func (h *VoyageHandler) GetVoyageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "voyage ID is required",
		})
	}

	voyage, err := h.voyageUseCase.GetVoyageByID(c.Context(), id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Failed to get voyage")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "voyage not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": voyage,
	})
}
