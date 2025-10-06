package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chats/sailing-backend/internal/config"
	"github.com/chats/sailing-backend/internal/delivery/http/handler"
	"github.com/chats/sailing-backend/internal/delivery/http/middleware"
	"github.com/chats/sailing-backend/internal/repository"
	"github.com/chats/sailing-backend/internal/usecase"
	"github.com/chats/sailing-backend/pkg/database"
	"github.com/chats/sailing-backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Initialize logger
	logger.InitLogger(cfg.LogLevel)

	log.Info().Msg("Starting Sailing Backend API...")

	// Connect to MongoDB
	db, err := database.ConnectMongoDB(cfg.MongoDBURI, cfg.MongoDBDatabase)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}

	// Initialize repositories
	voyageRepo := repository.NewVoyageRepository(db)
	checkpointRepo := repository.NewCheckpointRepository(db)
	gpsTrackRepo := repository.NewGPSTrackRepository(db)

	// Initialize use cases
	voyageUseCase := usecase.NewVoyageUseCase(voyageRepo)
	checkpointUseCase := usecase.NewCheckpointUseCase(checkpointRepo, voyageRepo)
	gpsTrackUseCase := usecase.NewGPSTrackUseCase(gpsTrackRepo, voyageRepo)

	// Initialize handlers
	voyageHandler := handler.NewVoyageHandler(voyageUseCase)
	checkpointHandler := handler.NewCheckpointHandler(checkpointUseCase)
	gpsTrackHandler := handler.NewGPSTrackHandler(gpsTrackUseCase)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
		AppName:      "Sailing Backend API",
	})

	// Setup security middlewares
	middleware.SetupSecurityMiddlewares(app)

	// Logger middleware
	app.Use(middleware.LoggerMiddleware())

	// Health check endpoint (no auth)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Sailing Backend API is running",
		})
	})

	// API v1 routes with authentication
	api := app.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())

	// Voyage routes
	api.Post("/voyages/depart", voyageHandler.Depart)
	api.Post("/voyages/arrive", voyageHandler.Arrive)
	api.Get("/voyages/all", voyageHandler.GetAllVoyages)
	api.Get("/voyage/:id", voyageHandler.GetVoyageByID)

	// Checkpoint routes
	api.Post("/checkpoints", checkpointHandler.CreateCheckpoint)
	api.Post("/checkpoints/batch", checkpointHandler.CreateCheckpointsBatch)

	// GPS Track routes
	api.Post("/gps-tracks", gpsTrackHandler.CreateGPSTrack)
	api.Post("/gps-tracks/batch", gpsTrackHandler.CreateGPSTracksBatch)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Info().Msg("Shutting down gracefully...")
		_ = app.Shutdown()
	}()

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Info().Str("port", cfg.Port).Msg("Server starting...")

	if err := app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

// customErrorHandler handles errors globally
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	log.Error().
		Err(err).
		Int("status", code).
		Str("path", c.Path()).
		Msg("Request error")

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
