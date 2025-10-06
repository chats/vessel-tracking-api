# Quick Reference - Sailing Backend API

## 🚀 Quick Start
```bash
# Start with Docker Compose
docker-compose up -d

# Or use the script
./start.sh

# Check health
curl http://localhost:8080/health
```

## 🔑 Authentication
```bash
# Using API Key (recommended for testing)
-H "X-API-Key: sailing-api-key-12345"

# Using JWT Bearer Token
-H "Authorization: Bearer <your-jwt-token>"
```

## 📍 Base URL
```
http://localhost:8080/api/v1
```

## 🛣️ Endpoints Summary

### Voyages
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/voyages/depart` | POST | Start new voyage |
| `/voyages/arrive` | POST | Complete voyage |
| `/voyages/all` | GET | List all voyages with checkpoints & GPS tracks |
| `/voyage/:id` | GET | Get voyage by ID |

### Checkpoints
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/checkpoints` | POST | Create checkpoint |
| `/checkpoints/batch` | POST | Bulk create checkpoints |

### GPS Tracks
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/gps-tracks` | POST | Create GPS track |
| `/gps-tracks/batch` | POST | Bulk create GPS tracks |

## 📝 Example Requests

### 1. Start Voyage
```bash
curl -X POST http://localhost:8080/api/v1/voyages/depart \
  -H "X-API-Key: sailing-api-key-12345" \
  -H "Content-Type: application/json" \
  -d '{
    "ship_id": "SHIP001",
    "ship_name": "Sea Explorer",
    "departure_port": "Bangkok Port"
  }'
```

### 2. Complete Voyage
```bash
curl -X POST http://localhost:8080/api/v1/voyages/arrive \
  -H "X-API-Key: sailing-api-key-12345" \
  -H "Content-Type: application/json" \
  -d '{
    "voyage_id": "<voyage-id>",
    "arrival_port": "Singapore Port"
  }'
```

### 3. List Voyages (with checkpoints and GPS tracks)
```bash
curl http://localhost:8080/api/v1/voyages/all?limit=10&offset=0 \
  -H "X-API-Key: sailing-api-key-12345"
```

**Response includes:**
- Voyage details
- All checkpoints for each voyage
- All GPS tracks for each voyage

### 4. Create Checkpoint
```bash
curl -X POST http://localhost:8080/api/v1/checkpoints \
  -H "X-API-Key: sailing-api-key-12345" \
  -H "Content-Type: application/json" \
  -d '{
    "voyage_id": "<voyage-id>",
    "location": {"latitude": 13.7563, "longitude": 100.5018},
    "timestamp": "2025-10-01T10:00:00Z",
    "description": "Checkpoint at Bangkok"
  }'
```

### 5. Create GPS Track
```bash
curl -X POST http://localhost:8080/api/v1/gps-tracks \
  -H "X-API-Key: sailing-api-key-12345" \
  -H "Content-Type: application/json" \
  -d '{
    "voyage_id": "<voyage-id>",
    "location": {"latitude": 13.7563, "longitude": 100.5018},
    "speed": 12.5,
    "heading": 90.0,
    "timestamp": "2025-10-01T10:00:00Z"
  }'
```

### 6. Batch Create Checkpoints
```bash
curl -X POST http://localhost:8080/api/v1/checkpoints/batch \
  -H "X-API-Key: sailing-api-key-12345" \
  -H "Content-Type: application/json" \
  -d '[
    {
      "voyage_id": "<voyage-id>",
      "location": {"latitude": 13.7563, "longitude": 100.5018},
      "timestamp": "2025-10-01T10:00:00Z"
    },
    {
      "voyage_id": "<voyage-id>",
      "location": {"latitude": 13.8563, "longitude": 100.6018},
      "timestamp": "2025-10-01T11:00:00Z"
    }
  ]'
```

## 🐳 Docker Commands
```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Restart services
docker-compose restart

# View API logs only
docker-compose logs -f api

# View MongoDB logs only
docker-compose logs -f mongodb
```

## 🔧 Development Commands
```bash
# Build
make build

# Run locally
make run

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean

# Install dev tools
make install-tools
```

## 📊 Response Formats

### Success Response
```json
{
  "message": "operation successful",
  "data": { /* result data */ }
}
```

### Error Response
```json
{
  "error": "error message"
}
```

### List Response
```json
{
  "data": [ /* array of items */ ],
  "count": 10
}
```

## ⚙️ Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| PORT | 8080 | Server port |
| ENV | development | Environment |
| MONGODB_URI | mongodb://localhost:27017 | MongoDB connection |
| MONGODB_DATABASE | sailing_db | Database name |
| JWT_SECRET | (required) | JWT secret key |
| API_KEY | (required) | API key |
| LOG_LEVEL | info | Log level |

## 🔍 Troubleshooting

### API not responding
```bash
# Check if services are running
docker-compose ps

# Check logs
docker-compose logs -f api

# Restart services
docker-compose restart
```

### MongoDB connection error
```bash
# Check MongoDB is running
docker-compose ps mongodb

# Check MongoDB logs
docker-compose logs mongodb

# Restart MongoDB
docker-compose restart mongodb
```

### Authentication error
- Verify API key in request header: `X-API-Key: sailing-api-key-12345`
- Check `.env` file for correct API_KEY value

## 📖 Documentation Files

- `README.md` - Full documentation
- `PROJECT_SUMMARY.md` - Project overview
- `QUICK_REFERENCE.md` - This file
- `api-tests.http` - API test examples

## 🎯 Common Use Cases

### Complete Voyage Flow
1. Start voyage (`/voyages/depart`)
2. Create GPS tracks as ship moves (`/gps-tracks`)
3. Create checkpoints at important locations (`/checkpoints`)
4. Complete voyage (`/voyages/arrive`)
5. View voyage history (`/voyages/all` or `/voyage/:id`)

### Batch Operations
- Use batch endpoints for efficiency when sending multiple records
- Recommended for GPS tracks collected over time
- Recommended for multiple checkpoints from the same voyage

## 🛡️ Security Best Practices

1. Change default API key in production
2. Use HTTPS in production
3. Rotate JWT secrets regularly
4. Monitor rate limit violations
5. Review access logs regularly
6. Keep dependencies updated

## 📞 Support

For detailed information, see:
- README.md for full documentation
- PROJECT_SUMMARY.md for architecture details
- api-tests.http for working examples
