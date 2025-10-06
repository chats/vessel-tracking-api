# Sailing Backend API - Architecture Diagram

## System Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLIENT LAYER                             │
│  (Mobile Apps, Web Apps, IoT Devices, External Services)        │
└────────────────┬────────────────────────────────────────────────┘
                 │
                 │ HTTP/HTTPS Requests
                 │ Authentication: JWT Bearer / API Key
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                    DELIVERY LAYER (HTTP)                         │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐    │
│  │   Middleware   │  │   Middleware   │  │   Middleware   │    │
│  │  Authentication│  │     Logger     │  │    Security    │    │
│  │  - JWT Token   │  │   - Zerolog    │  │  - CORS        │    │
│  │  - API Key     │  │   - Requests   │  │  - Helmet      │    │
│  └────────────────┘  └────────────────┘  │  - Rate Limit  │    │
│                                           │  - Recovery    │    │
│                                           └────────────────┘    │
│                                                                   │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐    │
│  │     Voyage     │  │   Checkpoint   │  │   GPS Track    │    │
│  │    Handler     │  │    Handler     │  │    Handler     │    │
│  │  - Depart      │  │  - Create      │  │  - Create      │    │
│  │  - Arrive      │  │  - Batch       │  │  - Batch       │    │
│  │  - Get All     │  │                │  │                │    │
│  │  - Get By ID   │  │                │  │                │    │
│  └────────┬───────┘  └────────┬───────┘  └────────┬───────┘    │
└───────────┼──────────────────┼──────────────────┼──────────────┘
            │                  │                  │
            │                  │                  │
            ▼                  ▼                  ▼
┌─────────────────────────────────────────────────────────────────┐
│                     USE CASE LAYER                               │
│                  (Business Logic)                                │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐    │
│  │     Voyage     │  │   Checkpoint   │  │   GPS Track    │    │
│  │    UseCase     │  │    UseCase     │  │    UseCase     │    │
│  │                │  │                │  │                │    │
│  │ - Validation   │  │ - Validation   │  │ - Validation   │    │
│  │ - Business     │  │ - Business     │  │ - Business     │    │
│  │   Rules        │  │   Rules        │  │   Rules        │    │
│  │ - Coordination │  │ - Coordination │  │ - Coordination │    │
│  └────────┬───────┘  └────────┬───────┘  └────────┬───────┘    │
└───────────┼──────────────────┼──────────────────┼──────────────┘
            │                  │                  │
            │                  │                  │
            ▼                  ▼                  ▼
┌─────────────────────────────────────────────────────────────────┐
│                    DOMAIN LAYER                                  │
│                  (Core Business)                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌────────────────────────────────────────────────────────┐     │
│  │              Domain Entities                           │     │
│  │  - Voyage      - Checkpoint    - GPSTrack             │     │
│  │  - Location    - WeatherInfo                          │     │
│  └────────────────────────────────────────────────────────┘     │
│                                                                   │
│  ┌────────────────────────────────────────────────────────┐     │
│  │           Repository Interfaces                        │     │
│  │  - VoyageRepository                                    │     │
│  │  - CheckpointRepository                                │     │
│  │  - GPSTrackRepository                                  │     │
│  └────────────────────────────────────────────────────────┘     │
└───────────────────────────┬─────────────────────────────────────┘
                            │
                            │ Implementation
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│                 REPOSITORY LAYER                                 │
│              (Data Access / MongoDB)                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐    │
│  │     Voyage     │  │   Checkpoint   │  │   GPS Track    │    │
│  │   Repository   │  │   Repository   │  │   Repository   │    │
│  │                │  │                │  │                │    │
│  │ - Create       │  │ - Create       │  │ - Create       │    │
│  │ - Update       │  │ - Batch Create │  │ - Batch Create │    │
│  │ - Find         │  │ - Find By      │  │ - Find By      │    │
│  │ - FindAll      │  │   Voyage       │  │   Voyage       │    │
│  └────────┬───────┘  └────────┬───────┘  └────────┬───────┘    │
└───────────┼──────────────────┼──────────────────┼──────────────┘
            │                  │                  │
            └──────────────────┴──────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                     DATABASE LAYER                               │
│                    MongoDB Database                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌──────────────────────────────────────────────────────┐       │
│  │  Collection: voyages                                 │       │
│  │  - Stores voyage information                         │       │
│  │  - Indexes: voyage_id, ship_id, departure_time       │       │
│  └──────────────────────────────────────────────────────┘       │
│                                                                   │
│  ┌──────────────────────────────────────────────────────┐       │
│  │  Collection: checkpoints                             │       │
│  │  - Stores checkpoint data with weather info          │       │
│  │  - Indexes: voyage_id, timestamp                     │       │
│  └──────────────────────────────────────────────────────┘       │
│                                                                   │
│  ┌──────────────────────────────────────────────────────┐       │
│  │  Collection: gps_tracks                              │       │
│  │  - Stores GPS tracking data                          │       │
│  │  - Indexes: voyage_id, timestamp                     │       │
│  └──────────────────────────────────────────────────────┘       │
└─────────────────────────────────────────────────────────────────┘
```

## Request Flow Example

### Example: Creating a New Voyage

```
1. Client Request
   POST /api/v1/voyages/depart
   Header: X-API-Key: sailing-api-key-12345
   Body: { ship_id, ship_name, departure_port }
   
   ↓
   
2. Security Middleware
   - CORS validation
   - Rate limit check
   - Panic recovery setup
   
   ↓
   
3. Auth Middleware
   - Validate API Key
   - OR Validate JWT Token
   - Return 401 if invalid
   
   ↓
   
4. Logger Middleware
   - Log request details
   - Track request timing
   
   ↓
   
5. Voyage Handler
   - Parse request body
   - Validate input
   - Call Use Case
   
   ↓
   
6. Voyage UseCase
   - Apply business rules
   - Validate ship_id, ship_name, port
   - Generate voyage_id (UUID)
   - Set status = "in_progress"
   - Set timestamps
   - Call Repository
   
   ↓
   
7. Voyage Repository
   - Map domain entity to MongoDB document
   - Execute InsertOne
   - Handle MongoDB errors
   - Return created entity
   
   ↓
   
8. MongoDB
   - Insert document
   - Apply indexes
   - Return inserted ID
   
   ↓
   
9. Response Flow (back up the stack)
   Repository → UseCase → Handler → Middleware → Client
   
   ↓
   
10. Client Response
    Status: 201 Created
    Body: {
      "message": "voyage departed successfully",
      "data": { voyage object with _id }
    }
```

## Component Dependencies

```
┌─────────────────────────────────────────────┐
│            External Dependencies             │
├─────────────────────────────────────────────┤
│  • github.com/gofiber/fiber/v2              │
│  • github.com/rs/zerolog                     │
│  • go.mongodb.org/mongo-driver               │
│  • github.com/golang-jwt/jwt/v5              │
│  • github.com/google/uuid                    │
│  • github.com/joho/godotenv                  │
└─────────────────────────────────────────────┘
```

## Clean Architecture Principles

```
┌────────────────────────────────────────────────┐
│         Dependency Rule                         │
│  (Dependencies point inward only)              │
└────────────────────────────────────────────────┘

        Delivery Layer
              ↓
         Use Case Layer
              ↓
         Domain Layer  ← Core (No dependencies)
              ↑
        Repository Layer
              ↓
         Database Layer
```

### Layer Independence

- **Domain Layer**: No external dependencies
- **Use Case Layer**: Depends only on Domain
- **Repository Layer**: Implements Domain interfaces
- **Delivery Layer**: Depends on Use Cases and Domain

## Security Architecture

```
┌─────────────────────────────────────────────────┐
│              Request Security                    │
├─────────────────────────────────────────────────┤
│                                                  │
│  1. Rate Limiting                               │
│     └─ 100 requests/minute per IP               │
│                                                  │
│  2. Authentication                               │
│     ├─ JWT Bearer Token                         │
│     └─ API Key Header                           │
│                                                  │
│  3. CORS                                         │
│     └─ Configurable origins                     │
│                                                  │
│  4. Helmet                                       │
│     └─ Security headers                         │
│                                                  │
│  5. Recovery                                     │
│     └─ Panic handling                           │
│                                                  │
│  6. Logging                                      │
│     └─ All requests logged                      │
│                                                  │
└─────────────────────────────────────────────────┘
```

## Data Flow Patterns

### Single Record Creation
```
Client → Handler → UseCase → Repository → MongoDB
                                             ↓
Client ← Handler ← UseCase ← Repository ← Success
```

### Batch Record Creation
```
Client → Handler → UseCase → Repository → MongoDB
(Array)                        (Array)      (InsertMany)
                                             ↓
Client ← Handler ← UseCase ← Repository ← Success
(Array)                        (Array)      (with IDs)
```

### Query Pattern
```
Client → Handler → UseCase → Repository → MongoDB
                                             ↓
                                           Find
                                             ↓
Client ← Handler ← UseCase ← Repository ← Documents
```

## Deployment Architecture

```
┌──────────────────────────────────────────┐
│         Docker Compose Setup              │
├──────────────────────────────────────────┤
│                                           │
│  ┌────────────────────────────────┐      │
│  │   API Container                │      │
│  │   - Port 8080                  │      │
│  │   - Golang Application         │      │
│  │   - Multi-stage build          │      │
│  └───────────┬────────────────────┘      │
│              │                            │
│              │ Network: sailing-network   │
│              │                            │
│  ┌───────────▼────────────────────┐      │
│  │   MongoDB Container            │      │
│  │   - Port 27017                 │      │
│  │   - Persistent Volume          │      │
│  │   - Init Script                │      │
│  └────────────────────────────────┘      │
│                                           │
└──────────────────────────────────────────┘
```

## Technology Stack Summary

| Layer | Technologies |
|-------|-------------|
| **Language** | Go 1.23 |
| **Web Framework** | Fiber v2 |
| **Logging** | Zerolog |
| **Database** | MongoDB 7.0 |
| **Authentication** | JWT, API Key |
| **Containerization** | Docker, Docker Compose |
| **Configuration** | Environment Variables |
| **Build Tool** | Make |

## File Organization

```
Clean Architecture Structure:

cmd/              → Application Entry Points
  └─ api/         → Main API application

internal/         → Private Application Code
  ├─ config/      → Configuration
  ├─ delivery/    → External Interfaces
  │   └─ http/    → HTTP Layer
  ├─ domain/      → Business Entities
  ├─ repository/  → Data Access
  └─ usecase/     → Business Logic

pkg/              → Public Libraries
  ├─ database/    → DB Connection
  └─ logger/      → Logging Setup
```

This architecture ensures:
- ✅ **Testability**: Each layer can be tested independently
- ✅ **Maintainability**: Clear separation of concerns
- ✅ **Scalability**: Easy to add new features
- ✅ **Flexibility**: Easy to swap implementations
- ✅ **Independence**: Business logic independent of frameworks
