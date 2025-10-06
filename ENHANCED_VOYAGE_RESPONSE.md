# GET /api/v1/voyages/all - Enhanced Response Example

## Overview

The endpoint `/api/v1/voyages/all` now returns **complete voyage details** including:
- Full voyage information
- All checkpoints associated with each voyage
- All GPS tracks associated with each voyage

This provides a comprehensive view of each voyage's journey in a single API call.

---

## Request

```bash
curl http://localhost:8080/api/v1/voyages/all?limit=10&offset=0 \
  -H "X-API-Key: sailing-api-key-12345"
```

### Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `limit` | integer | 100 | Maximum number of voyages to return |
| `offset` | integer | 0 | Number of voyages to skip (for pagination) |

---

## Response Structure

```json
{
  "data": [
    {
      "voyage": { /* Voyage object */ },
      "checkpoints": [ /* Array of Checkpoint objects */ ],
      "gps_tracks": [ /* Array of GPSTrack objects */ ]
    }
  ],
  "count": 1
}
```

---

## Complete Response Example

```json
{
  "data": [
    {
      "voyage": {
        "id": "6707a1b2c3d4e5f6a7b8c9d0",
        "voyage_id": "voyage-bangkok-singapore-001",
        "ship_id": "SHIP001",
        "ship_name": "Sea Explorer",
        "departure_port": "Bangkok Port",
        "arrival_port": "Singapore Port",
        "departure_time": "2025-10-06T08:00:00Z",
        "arrival_time": "2025-10-07T20:00:00Z",
        "status": "completed",
        "created_at": "2025-10-06T08:00:00Z",
        "updated_at": "2025-10-07T20:00:00Z"
      },
      "checkpoints": [
        {
          "id": "6707a2b3c4d5e6f7a8b9c0d1",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 13.7563,
            "longitude": 100.5018
          },
          "timestamp": "2025-10-06T08:00:00Z",
          "description": "Departure from Bangkok Port",
          "weather": {
            "temperature": 30.0,
            "wind_speed": 10.0,
            "wind_dir": 180.0,
            "wave_height": 0.8,
            "condition": "clear"
          },
          "created_at": "2025-10-06T08:00:00Z"
        },
        {
          "id": "6707a3b4c5d6e7f8a9b0c1d2",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 12.5657,
            "longitude": 99.9576
          },
          "timestamp": "2025-10-06T12:00:00Z",
          "description": "Passing Hua Hin",
          "weather": {
            "temperature": 31.5,
            "wind_speed": 12.0,
            "wind_dir": 170.0,
            "wave_height": 1.0,
            "condition": "clear"
          },
          "created_at": "2025-10-06T12:00:00Z"
        },
        {
          "id": "6707a4b5c6d7e8f9a0b1c2d3",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 7.8804,
            "longitude": 98.3923
          },
          "timestamp": "2025-10-07T02:00:00Z",
          "description": "Phuket Waters",
          "weather": {
            "temperature": 28.5,
            "wind_speed": 15.0,
            "wind_dir": 220.0,
            "wave_height": 1.2,
            "condition": "partly cloudy"
          },
          "created_at": "2025-10-07T02:00:00Z"
        },
        {
          "id": "6707a5b6c7d8e9f0a1b2c3d4",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 1.2521,
            "longitude": 103.8198
          },
          "timestamp": "2025-10-07T14:00:00Z",
          "description": "Arrival at Singapore Strait",
          "weather": {
            "temperature": 30.0,
            "wind_speed": 10.5,
            "wind_dir": 90.0,
            "wave_height": 0.7,
            "condition": "clear"
          },
          "created_at": "2025-10-07T14:00:00Z"
        }
      ],
      "gps_tracks": [
        {
          "id": "6707a6b7c8d9e0f1a2b3c4d5",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 13.7563,
            "longitude": 100.5018
          },
          "speed": 12.5,
          "heading": 180.0,
          "altitude": 5.0,
          "timestamp": "2025-10-06T08:00:00Z",
          "created_at": "2025-10-06T08:00:00Z"
        },
        {
          "id": "6707a7b8c9d0e1f2a3b4c5d6",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 13.7463,
            "longitude": 100.4918
          },
          "speed": 13.2,
          "heading": 181.5,
          "altitude": 5.2,
          "timestamp": "2025-10-06T08:05:00Z",
          "created_at": "2025-10-06T08:05:00Z"
        },
        {
          "id": "6707a8b9c0d1e2f3a4b5c6d7",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 13.7363,
            "longitude": 100.4818
          },
          "speed": 13.8,
          "heading": 183.0,
          "altitude": 5.5,
          "timestamp": "2025-10-06T08:10:00Z",
          "created_at": "2025-10-06T08:10:00Z"
        },
        {
          "id": "6707a9b0c1d2e3f4a5b6c7d8",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 12.5657,
            "longitude": 99.9576
          },
          "speed": 14.5,
          "heading": 185.0,
          "altitude": 6.0,
          "timestamp": "2025-10-06T12:00:00Z",
          "created_at": "2025-10-06T12:00:00Z"
        },
        {
          "id": "6707aab1c2d3e4f5a6b7c8d9",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 7.8804,
            "longitude": 98.3923
          },
          "speed": 15.2,
          "heading": 190.0,
          "altitude": 7.0,
          "timestamp": "2025-10-07T02:00:00Z",
          "created_at": "2025-10-07T02:00:00Z"
        },
        {
          "id": "6707abb2c3d4e5f6a7b8c9da",
          "voyage_id": "voyage-bangkok-singapore-001",
          "location": {
            "latitude": 1.2521,
            "longitude": 103.8198
          },
          "speed": 3.5,
          "heading": 45.0,
          "altitude": 2.5,
          "timestamp": "2025-10-07T19:50:00Z",
          "created_at": "2025-10-07T19:50:00Z"
        }
      ]
    },
    {
      "voyage": {
        "id": "6707acb3c4d5e6f7a8b9c0db",
        "voyage_id": "voyage-phuket-hongkong-002",
        "ship_id": "SHIP002",
        "ship_name": "Pacific Trader",
        "departure_port": "Phuket Port",
        "arrival_port": "",
        "departure_time": "2025-10-07T10:00:00Z",
        "arrival_time": null,
        "status": "in_progress",
        "created_at": "2025-10-07T10:00:00Z",
        "updated_at": "2025-10-07T10:00:00Z"
      },
      "checkpoints": [
        {
          "id": "6707adb4c5d6e7f8a9b0c1dc",
          "voyage_id": "voyage-phuket-hongkong-002",
          "location": {
            "latitude": 7.8804,
            "longitude": 98.3923
          },
          "timestamp": "2025-10-07T10:00:00Z",
          "description": "Departure from Phuket",
          "weather": {
            "temperature": 29.0,
            "wind_speed": 12.5,
            "wind_dir": 180.0,
            "wave_height": 1.0,
            "condition": "partly cloudy"
          },
          "created_at": "2025-10-07T10:00:00Z"
        }
      ],
      "gps_tracks": [
        {
          "id": "6707aeb5c6d7e8f9a0b1c2dd",
          "voyage_id": "voyage-phuket-hongkong-002",
          "location": {
            "latitude": 7.8804,
            "longitude": 98.3923
          },
          "speed": 15.0,
          "heading": 90.0,
          "altitude": 8.0,
          "timestamp": "2025-10-07T10:00:00Z",
          "created_at": "2025-10-07T10:00:00Z"
        },
        {
          "id": "6707afb6c7d8e9f0a1b2c3de",
          "voyage_id": "voyage-phuket-hongkong-002",
          "location": {
            "latitude": 7.8904,
            "longitude": 98.4023
          },
          "speed": 15.5,
          "heading": 91.0,
          "altitude": 8.2,
          "timestamp": "2025-10-07T10:05:00Z",
          "created_at": "2025-10-07T10:05:00Z"
        }
      ]
    }
  ],
  "count": 2
}
```

---

## Response Fields Explanation

### Voyage Object
- `id` - MongoDB ObjectID
- `voyage_id` - Unique voyage identifier (UUID or custom)
- `ship_id` - Ship identifier
- `ship_name` - Name of the ship
- `departure_port` - Port of departure
- `arrival_port` - Port of arrival (empty if in progress)
- `departure_time` - ISO 8601 timestamp of departure
- `arrival_time` - ISO 8601 timestamp of arrival (null if in progress)
- `status` - Voyage status: `in_progress`, `completed`, or `cancelled`
- `created_at` - Record creation timestamp
- `updated_at` - Record last update timestamp

### Checkpoints Array
Contains all checkpoints recorded during the voyage:
- `id` - MongoDB ObjectID
- `voyage_id` - Reference to the voyage
- `location` - GPS coordinates (latitude, longitude)
- `timestamp` - When the checkpoint was recorded
- `description` - Optional description
- `weather` - Optional weather information
  - `temperature` - In Celsius
  - `wind_speed` - In knots
  - `wind_dir` - Wind direction in degrees
  - `wave_height` - In meters
  - `condition` - Weather condition text
- `created_at` - Record creation timestamp

### GPS Tracks Array
Contains all GPS tracking points during the voyage:
- `id` - MongoDB ObjectID
- `voyage_id` - Reference to the voyage
- `location` - GPS coordinates (latitude, longitude)
- `speed` - Ship speed in knots
- `heading` - Direction in degrees (0-360)
- `altitude` - Altitude/elevation in meters
- `timestamp` - When the GPS data was recorded
- `created_at` - Record creation timestamp

---

## Use Cases

### 1. Voyage History and Analysis
Get complete journey information including:
- Route taken (from GPS tracks)
- Important waypoints (from checkpoints)
- Weather conditions during journey
- Speed and heading changes

### 2. Voyage Monitoring Dashboard
Display real-time and historical data:
- Current voyage status
- All tracked positions
- Weather history
- Speed profile

### 3. Route Optimization
Analyze past voyages to:
- Compare different routes
- Evaluate weather impact
- Calculate average speeds
- Identify optimal checkpoints

### 4. Reporting and Compliance
Generate comprehensive reports:
- Complete voyage logs
- Weather encounter reports
- Speed and position history
- Checkpoint verification

---

## Pagination

Use `limit` and `offset` for pagination:

```bash
# First page (10 items)
curl http://localhost:8080/api/v1/voyages/all?limit=10&offset=0 \
  -H "X-API-Key: sailing-api-key-12345"

# Second page (10 items)
curl http://localhost:8080/api/v1/voyages/all?limit=10&offset=10 \
  -H "X-API-Key: sailing-api-key-12345"

# Third page (10 items)
curl http://localhost:8080/api/v1/voyages/all?limit=10&offset=20 \
  -H "X-API-Key: sailing-api-key-12345"
```

---

## Performance Notes

- Each voyage includes **all** its checkpoints and GPS tracks
- For voyages with many GPS tracks, the response can be large
- Use pagination (`limit`) to control response size
- Consider implementing filtering options in future versions
- GPS tracks are typically recorded every 5 minutes, resulting in ~288 tracks per day
- Checkpoints are usually fewer (4-10 per voyage)

---

## Example Response Sizes

| Voyage Duration | Estimated GPS Tracks | Estimated Checkpoints | Approx Response Size |
|-----------------|---------------------|----------------------|---------------------|
| 6 hours | ~72 tracks | 3-5 checkpoints | ~50 KB |
| 12 hours | ~144 tracks | 5-8 checkpoints | ~90 KB |
| 24 hours | ~288 tracks | 8-12 checkpoints | ~180 KB |
| 48 hours | ~576 tracks | 12-20 checkpoints | ~350 KB |

---

## Benefits of Enhanced Response

### ✅ Single API Call
- No need to make multiple requests
- Reduces network overhead
- Simplifies client code

### ✅ Complete Journey View
- All voyage data in one response
- Easy to visualize on maps
- Comprehensive analysis possible

### ✅ Better Performance for Dashboards
- Load all data at once
- Faster initial rendering
- Better user experience

### ✅ Simplified Data Management
- Consistent data structure
- Easy to cache
- Straightforward to process

---

## Alternative Approaches (Future Enhancements)

### Option 1: Filtered Response
```bash
# Only include checkpoints
GET /api/v1/voyages/all?include=checkpoints

# Only include GPS tracks
GET /api/v1/voyages/all?include=gps_tracks

# Include both (default)
GET /api/v1/voyages/all?include=checkpoints,gps_tracks
```

### Option 2: Time-based Filtering
```bash
# Get voyages from last 7 days with all data
GET /api/v1/voyages/all?since=7d

# Get voyages between dates
GET /api/v1/voyages/all?from=2025-10-01&to=2025-10-07
```

### Option 3: Summary vs Full Mode
```bash
# Summary mode (voyage info only)
GET /api/v1/voyages/all?mode=summary

# Full mode (with checkpoints and GPS tracks)
GET /api/v1/voyages/all?mode=full
```

---

## Error Handling

If there's an error fetching checkpoints or GPS tracks for a specific voyage:
- The response will include **empty arrays** for those fields
- The voyage data will still be returned
- No error is thrown, ensuring partial data is available

Example with missing data:
```json
{
  "voyage": { /* voyage data */ },
  "checkpoints": [],  // Empty if error occurred
  "gps_tracks": []    // Empty if error occurred
}
```

This ensures the API is resilient and always returns useful data.
