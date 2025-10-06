# Sailing Backend API

A RESTful API for sailing/voyage tracking system built with Go, Fiber, MongoDB following Clean Architecture principles.

## Features

- ✅ Clean Architecture (Domain, Use Case, Repository, Delivery layers)
- ✅ Golang with Fiber web framework
- ✅ Structured logging with Zerolog
- ✅ Authentication with JWT Bearer tokens and API Keys
- ✅ MongoDB for data persistence
- ✅ Security middlewares (CORS, Helmet, Rate Limiting, Recovery)
- ✅ Docker and Docker Compose support
- ✅ Graceful shutdown

## Architecture

```
sailing-backend/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── delivery/
│   │   └── http/
│   │       ├── handler/            # HTTP handlers
│   │       │   ├── voyage_handler.go
│   │       │   ├── checkpoint_handler.go
│   │       │   └── gps_track_handler.go
│   │       └── middleware/         # HTTP middlewares
│   │           ├── auth.go
│   │           ├── logger.go
│   │           └── security.go
│   ├── domain/
│   │   ├── entities.go             # Domain entities
│   │   └── repository.go           # Repository interfaces
│   ├── repository/                 # Repository implementations
│   │   ├── voyage_repository.go
│   │   ├── checkpoint_repository.go
│   │   └── gps_track_repository.go
│   └── usecase/                    # Business logic
│       ├── voyage_usecase.go
│       ├── checkpoint_usecase.go
│       └── gps_track_usecase.go
├── pkg/
│   ├── database/
│   │   └── mongodb.go              # Database connection
│   └── logger/
│       └── logger.go               # Logger configuration
├── .env.example                    # Environment variables template
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
└── init-mongo.js                   # MongoDB initialization script
```

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint (no authentication required)

### Voyage Management
- `POST /api/v1/voyages/depart` - Create a new voyage (departure)
- `POST /api/v1/voyages/arrive` - Update voyage with arrival information
- `GET /api/v1/voyages/all` - Get all voyages (with pagination)
- `GET /api/v1/voyage/:id` - Get voyage by ID

### Checkpoint Management
- `POST /api/v1/checkpoints` - Create a single checkpoint
- `POST /api/v1/checkpoints/batch` - Create multiple checkpoints

### GPS Track Management
- `POST /api/v1/gps-tracks` - Create a single GPS track
- `POST /api/v1/gps-tracks/batch` - Create multiple GPS tracks

## Authentication

The API supports two authentication methods:

### 1. Bearer Token (JWT)
```bash
Authorization: Bearer <your-jwt-token>
```

### 2. API Key
```bash
X-API-Key: <your-api-key>
```

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose
- MongoDB (if running locally without Docker)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd sailing-backend
```

2. Copy the environment file:
```bash
cp .env.example .env
```

3. Update the `.env` file with your configuration:
```env
PORT=8080
ENV=development
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=sailing_db
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
API_KEY=your-api-key-change-this-in-production
LOG_LEVEL=info
```

### Running with Docker Compose

The easiest way to run the application:

```bash
# Build and start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

This will start:
- MongoDB on port 27017
- API server on port 8080

### Running Locally

1. Install dependencies:
```bash
go mod download
```

2. Start MongoDB (if not using Docker):
```bash
# Using Docker for MongoDB only
docker run -d -p 27017:27017 --name mongodb mongo:7.0
```

3. Run the application:
```bash
go run cmd/api/main.go
```

### Building the Application

```bash
# Build binary
go build -o bin/api cmd/api/main.go

# Run binary
./bin/api
```

## API Usage Examples

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. Depart Voyage
```bash
curl -X POST http://localhost:8080/api/v1/voyages/depart \
  -H "X-API-Key: your-api-key-change-this-in-production" \
  -H "Content-Type: application/json" \
  -d '{
    "ship_id": "SHIP001",
    "ship_name": "Sea Explorer",
    "departure_port": "Bangkok Port"
  }'
```

### 3. Arrive Voyage
```bash
curl -X POST http://localhost:8080/api/v1/voyages/arrive \
  -H "X-API-Key: your-api-key-change-this-in-production" \
  -H "Content-Type: application/json" \
  -d '{
    "voyage_id": "voyage-uuid",
    "arrival_port": "Singapore Port"
  }'
```

### 4. Create Checkpoint
```bash
curl -X POST http://localhost:8080/api/v1/checkpoints \
  -H "X-API-Key: your-api-key-change-this-in-production" \
  -H "Content-Type: application/json" \
  -d '{
    "voyage_id": "voyage-uuid",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "timestamp": "2025-10-01T10:00:00Z",
    "description": "Checkpoint at Bangkok",
    "weather": {
      "temperature": 30.5,
      "wind_speed": 15.0,
      "wind_dir": 180.0,
      "wave_height": 1.5,
      "condition": "clear"
    }
  }'
```

### 5. Create GPS Track
```bash
curl -X POST http://localhost:8080/api/v1/gps-tracks \
  -H "X-API-Key: your-api-key-change-this-in-production" \
  -H "Content-Type: application/json" \
  -d '{
    "voyage_id": "voyage-uuid",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "speed": 12.5,
    "heading": 90.0,
    "altitude": 10.0,
    "timestamp": "2025-10-01T10:00:00Z"
  }'
```

### 6. Get All Voyages
```bash
curl http://localhost:8080/api/v1/voyages/all?limit=10&offset=0 \
  -H "X-API-Key: your-api-key-change-this-in-production"
```

### 7. Get Voyage by ID
```bash
curl http://localhost:8080/api/v1/voyage/67890abcdef \
  -H "X-API-Key: your-api-key-change-this-in-production"
```

### 8. Batch Create Checkpoints
```bash
curl -X POST http://localhost:8080/api/v1/checkpoints/batch \
  -H "X-API-Key: your-api-key-change-this-in-production" \
  -H "Content-Type: application/json" \
  -d '[
    {
      "voyage_id": "voyage-uuid",
      "location": {"latitude": 13.7563, "longitude": 100.5018},
      "timestamp": "2025-10-01T10:00:00Z",
      "description": "Checkpoint 1"
    },
    {
      "voyage_id": "voyage-uuid",
      "location": {"latitude": 13.8563, "longitude": 100.6018},
      "timestamp": "2025-10-01T11:00:00Z",
      "description": "Checkpoint 2"
    }
  ]'
```

## Data Models

### Voyage
```json
{
  "id": "ObjectID",
  "voyage_id": "unique-voyage-id",
  "ship_id": "SHIP001",
  "ship_name": "Sea Explorer",
  "departure_port": "Bangkok Port",
  "arrival_port": "Singapore Port",
  "departure_time": "2025-10-01T08:00:00Z",
  "arrival_time": "2025-10-02T20:00:00Z",
  "status": "completed",
  "created_at": "2025-10-01T08:00:00Z",
  "updated_at": "2025-10-02T20:00:00Z"
}
```

### Checkpoint
```json
{
  "id": "ObjectID",
  "voyage_id": "voyage-uuid",
  "location": {
    "latitude": 13.7563,
    "longitude": 100.5018
  },
  "timestamp": "2025-10-01T10:00:00Z",
  "description": "Checkpoint description",
  "weather": {
    "temperature": 30.5,
    "wind_speed": 15.0,
    "wind_dir": 180.0,
    "wave_height": 1.5,
    "condition": "clear"
  },
  "created_at": "2025-10-01T10:00:00Z"
}
```

### GPS Track
```json
{
  "id": "ObjectID",
  "voyage_id": "voyage-uuid",
  "location": {
    "latitude": 13.7563,
    "longitude": 100.5018
  },
  "speed": 12.5,
  "heading": 90.0,
  "altitude": 10.0,
  "timestamp": "2025-10-01T10:00:00Z",
  "created_at": "2025-10-01T10:00:00Z"
}
```

## Security Features

- **CORS**: Cross-Origin Resource Sharing enabled
- **Helmet**: Security headers set automatically
- **Rate Limiting**: 100 requests per minute per IP
- **Recovery**: Automatic panic recovery
- **Authentication**: JWT Bearer tokens and API Key support

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 8080 |
| ENV | Environment (development/production) | development |
| MONGODB_URI | MongoDB connection string | mongodb://localhost:27017 |
| MONGODB_DATABASE | MongoDB database name | sailing_db |
| JWT_SECRET | JWT secret key | (change in production) |
| API_KEY | API key for authentication | (change in production) |
| LOG_LEVEL | Log level (debug/info/warn/error) | info |

## Development

### Running Tests
```bash
go test ./...
```

### Code Formatting
```bash
go fmt ./...
```

### Code Linting
```bash
golangci-lint run
```

## Production Deployment

1. Update environment variables in `.env` or `docker-compose.yaml`
2. Change default JWT_SECRET and API_KEY
3. Set ENV=production
4. Use strong MongoDB credentials
5. Configure proper MongoDB replica set for production
6. Set up proper logging and monitoring
7. Use HTTPS/TLS for secure communication

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Support

For issues and questions, please open an issue in the repository.
