# Sample POST Request Bodies - Sailing Backend API

This file contains example JSON bodies for all POST endpoints in the Sailing Backend API.

---

## 1. Voyage Depart - `/api/v1/voyages/depart`

### Minimal Request (Required fields only)
```json
{
  "ship_id": "SHIP001",
  "ship_name": "Sea Explorer",
  "departure_port": "Bangkok Port"
}
```

### Full Request (With optional voyage_id)
```json
{
  "voyage_id": "custom-voyage-id-12345",
  "ship_id": "SHIP001",
  "ship_name": "Sea Explorer",
  "departure_port": "Bangkok Port"
}
```

### Example with Thai Ports
```json
{
  "ship_id": "SHIP-TH-001",
  "ship_name": "ไทยสมุทร",
  "departure_port": "ท่าเรือแหลมฉบัง"
}
```

### Example - Container Ship
```json
{
  "ship_id": "CONTAINER-2024-001",
  "ship_name": "Pacific Carrier",
  "departure_port": "Port of Singapore"
}
```

### Example - Fishing Vessel
```json
{
  "ship_id": "FISH-099",
  "ship_name": "Ocean Hunter",
  "departure_port": "Phuket Deep Sea Port"
}
```

---

## 2. Voyage Arrive - `/api/v1/voyages/arrive`

### Basic Arrival
```json
{
  "voyage_id": "550e8400-e29b-41d4-a716-446655440000",
  "arrival_port": "Singapore Port"
}
```

### Example - Bangkok to Singapore
```json
{
  "voyage_id": "voyage-001-2025",
  "arrival_port": "Port of Singapore"
}
```

### Example - Domestic Route
```json
{
  "voyage_id": "domestic-voyage-123",
  "arrival_port": "ท่าเรือสงขลา"
}
```

### Example - International Route
```json
{
  "voyage_id": "intl-20250101-001",
  "arrival_port": "Port of Hong Kong"
}
```

---

## 3. Create Single Checkpoint - `/api/v1/checkpoints`

### Minimal Checkpoint (Required fields only)
```json
{
  "voyage_id": "550e8400-e29b-41d4-a716-446655440000",
  "location": {
    "latitude": 13.7563,
    "longitude": 100.5018
  },
  "timestamp": "2025-10-06T10:00:00Z"
}
```

### Checkpoint with Description
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 13.7563,
    "longitude": 100.5018
  },
  "timestamp": "2025-10-06T10:00:00Z",
  "description": "Departure from Bangkok Port"
}
```

### Checkpoint with Full Weather Data
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 7.8804,
    "longitude": 98.3923
  },
  "timestamp": "2025-10-06T14:30:00Z",
  "description": "Passing Phuket waters",
  "weather": {
    "temperature": 32.5,
    "wind_speed": 15.5,
    "wind_dir": 180.0,
    "wave_height": 1.2,
    "condition": "clear"
  }
}
```

### Checkpoint - Rough Sea Conditions
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 6.9271,
    "longitude": 100.3689
  },
  "timestamp": "2025-10-06T18:00:00Z",
  "description": "Andaman Sea - Rough conditions",
  "weather": {
    "temperature": 29.8,
    "wind_speed": 35.0,
    "wind_dir": 270.0,
    "wave_height": 3.5,
    "condition": "stormy"
  }
}
```

### Checkpoint - Calm Waters
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 1.2521,
    "longitude": 103.8198
  },
  "timestamp": "2025-10-07T08:00:00Z",
  "description": "Approaching Singapore Strait",
  "weather": {
    "temperature": 28.0,
    "wind_speed": 8.5,
    "wind_dir": 90.0,
    "wave_height": 0.5,
    "condition": "partly cloudy"
  }
}
```

---

## 4. Create Checkpoints Batch - `/api/v1/checkpoints/batch`

### Minimal Batch (2 checkpoints)
```json
[
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "timestamp": "2025-10-06T10:00:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 12.9236,
      "longitude": 100.8825
    },
    "timestamp": "2025-10-06T12:00:00Z"
  }
]
```

### Complete Journey Checkpoints (5 points)
```json
[
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "timestamp": "2025-10-06T08:00:00Z",
    "description": "Departure - Bangkok Port",
    "weather": {
      "temperature": 30.0,
      "wind_speed": 10.0,
      "wind_dir": 180.0,
      "wave_height": 0.8,
      "condition": "clear"
    }
  },
  {
    "voyage_id": "voyage-001-2025",
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
    }
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 9.1382,
      "longitude": 99.3330
    },
    "timestamp": "2025-10-06T18:00:00Z",
    "description": "Approaching Surat Thani",
    "weather": {
      "temperature": 29.0,
      "wind_speed": 18.0,
      "wind_dir": 200.0,
      "wave_height": 1.5,
      "condition": "cloudy"
    }
  },
  {
    "voyage_id": "voyage-001-2025",
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
    }
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 1.2521,
      "longitude": 103.8198
    },
    "timestamp": "2025-10-07T14:00:00Z",
    "description": "Arrival - Singapore Strait",
    "weather": {
      "temperature": 30.0,
      "wind_speed": 10.5,
      "wind_dir": 90.0,
      "wave_height": 0.7,
      "condition": "clear"
    }
  }
]
```

### Hourly Checkpoints (3 hours)
```json
[
  {
    "voyage_id": "voyage-002-2025",
    "location": {
      "latitude": 13.0000,
      "longitude": 100.5000
    },
    "timestamp": "2025-10-06T10:00:00Z",
    "description": "Hour 1"
  },
  {
    "voyage_id": "voyage-002-2025",
    "location": {
      "latitude": 12.9000,
      "longitude": 100.4000
    },
    "timestamp": "2025-10-06T11:00:00Z",
    "description": "Hour 2"
  },
  {
    "voyage_id": "voyage-002-2025",
    "location": {
      "latitude": 12.8000,
      "longitude": 100.3000
    },
    "timestamp": "2025-10-06T12:00:00Z",
    "description": "Hour 3"
  }
]
```

---

## 5. Create Single GPS Track - `/api/v1/gps-tracks`

### Minimal GPS Track (Required fields only)
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 13.7563,
    "longitude": 100.5018
  },
  "speed": 12.5,
  "heading": 180.0,
  "timestamp": "2025-10-06T10:00:00Z"
}
```

### GPS Track with Altitude
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 13.7563,
    "longitude": 100.5018
  },
  "speed": 15.8,
  "heading": 185.5,
  "altitude": 5.2,
  "timestamp": "2025-10-06T10:00:00Z"
}
```

### GPS Track - High Speed
```json
{
  "voyage_id": "cargo-express-001",
  "location": {
    "latitude": 7.8804,
    "longitude": 98.3923
  },
  "speed": 22.5,
  "heading": 270.0,
  "altitude": 8.0,
  "timestamp": "2025-10-06T15:30:00Z"
}
```

### GPS Track - Low Speed (Docking)
```json
{
  "voyage_id": "voyage-001-2025",
  "location": {
    "latitude": 1.2521,
    "longitude": 103.8198
  },
  "speed": 3.2,
  "heading": 45.0,
  "altitude": 2.5,
  "timestamp": "2025-10-07T14:50:00Z"
}
```

### GPS Track - Stationary (Anchored)
```json
{
  "voyage_id": "fishing-boat-099",
  "location": {
    "latitude": 8.5241,
    "longitude": 99.8258
  },
  "speed": 0.0,
  "heading": 0.0,
  "altitude": 0.0,
  "timestamp": "2025-10-06T20:00:00Z"
}
```

---

## 6. Create GPS Tracks Batch - `/api/v1/gps-tracks/batch`

### Minimal Batch (2 tracks)
```json
[
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "speed": 12.5,
    "heading": 180.0,
    "timestamp": "2025-10-06T10:00:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7463,
      "longitude": 100.4918
    },
    "speed": 13.0,
    "heading": 182.0,
    "timestamp": "2025-10-06T10:05:00Z"
  }
]
```

### 5-Minute Interval Tracking (30 minutes)
```json
[
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "speed": 12.5,
    "heading": 180.0,
    "altitude": 5.0,
    "timestamp": "2025-10-06T10:00:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7463,
      "longitude": 100.4918
    },
    "speed": 13.2,
    "heading": 181.5,
    "altitude": 5.2,
    "timestamp": "2025-10-06T10:05:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7363,
      "longitude": 100.4818
    },
    "speed": 13.8,
    "heading": 183.0,
    "altitude": 5.5,
    "timestamp": "2025-10-06T10:10:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7263,
      "longitude": 100.4718
    },
    "speed": 14.5,
    "heading": 184.5,
    "altitude": 6.0,
    "timestamp": "2025-10-06T10:15:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7163,
      "longitude": 100.4618
    },
    "speed": 15.0,
    "heading": 186.0,
    "altitude": 6.2,
    "timestamp": "2025-10-06T10:20:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.7063,
      "longitude": 100.4518
    },
    "speed": 15.2,
    "heading": 187.0,
    "altitude": 6.5,
    "timestamp": "2025-10-06T10:25:00Z"
  },
  {
    "voyage_id": "voyage-001-2025",
    "location": {
      "latitude": 13.6963,
      "longitude": 100.4418
    },
    "speed": 15.5,
    "heading": 188.5,
    "altitude": 7.0,
    "timestamp": "2025-10-06T10:30:00Z"
  }
]
```

### Acceleration Pattern (Speed increasing)
```json
[
  {
    "voyage_id": "speedboat-001",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5018
    },
    "speed": 5.0,
    "heading": 90.0,
    "altitude": 2.0,
    "timestamp": "2025-10-06T11:00:00Z"
  },
  {
    "voyage_id": "speedboat-001",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5118
    },
    "speed": 10.0,
    "heading": 90.0,
    "altitude": 3.0,
    "timestamp": "2025-10-06T11:02:00Z"
  },
  {
    "voyage_id": "speedboat-001",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5218
    },
    "speed": 15.0,
    "heading": 90.0,
    "altitude": 4.0,
    "timestamp": "2025-10-06T11:04:00Z"
  },
  {
    "voyage_id": "speedboat-001",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5318
    },
    "speed": 20.0,
    "heading": 90.0,
    "altitude": 5.0,
    "timestamp": "2025-10-06T11:06:00Z"
  },
  {
    "voyage_id": "speedboat-001",
    "location": {
      "latitude": 13.7563,
      "longitude": 100.5418
    },
    "speed": 25.0,
    "heading": 90.0,
    "altitude": 6.0,
    "timestamp": "2025-10-06T11:08:00Z"
  }
]
```

### Course Change Pattern (Turning)
```json
[
  {
    "voyage_id": "tanker-001",
    "location": {
      "latitude": 13.0000,
      "longitude": 100.0000
    },
    "speed": 15.0,
    "heading": 0.0,
    "altitude": 10.0,
    "timestamp": "2025-10-06T12:00:00Z"
  },
  {
    "voyage_id": "tanker-001",
    "location": {
      "latitude": 13.0100,
      "longitude": 100.0050
    },
    "speed": 14.5,
    "heading": 45.0,
    "altitude": 10.0,
    "timestamp": "2025-10-06T12:05:00Z"
  },
  {
    "voyage_id": "tanker-001",
    "location": {
      "latitude": 13.0150,
      "longitude": 100.0150
    },
    "speed": 14.0,
    "heading": 90.0,
    "altitude": 10.0,
    "timestamp": "2025-10-06T12:10:00Z"
  },
  {
    "voyage_id": "tanker-001",
    "location": {
      "latitude": 13.0150,
      "longitude": 100.0250
    },
    "speed": 14.5,
    "heading": 90.0,
    "altitude": 10.0,
    "timestamp": "2025-10-06T12:15:00Z"
  }
]
```

---

## Notes on Field Values

### Coordinates (Thailand & Southeast Asia)
- **Bangkok**: 13.7563°N, 100.5018°E
- **Phuket**: 7.8804°N, 98.3923°E
- **Singapore**: 1.2521°N, 103.8198°E
- **Hong Kong**: 22.3193°N, 114.1694°E

### Speed (in knots)
- **Slow/Docking**: 0-5 knots
- **Normal Cruising**: 10-15 knots
- **Fast Cruising**: 15-20 knots
- **High Speed**: 20-30 knots

### Heading (in degrees)
- **North**: 0° or 360°
- **East**: 90°
- **South**: 180°
- **West**: 270°

### Weather Conditions
- **Clear**: "clear"
- **Partly Cloudy**: "partly cloudy"
- **Cloudy**: "cloudy"
- **Rainy**: "rainy"
- **Stormy**: "stormy"

### Temperature
- Celsius (°C)
- Typical range: 25-35°C for Southeast Asia

### Wind Speed
- Knots
- Typical range: 5-40 knots

### Wave Height
- Meters
- Typical range: 0.5-5.0 meters

---

## Tips for Testing

1. **Use realistic coordinates** for your region
2. **Timestamps should be sequential** for GPS tracks and checkpoints
3. **Speed and heading should be consistent** with coordinate changes
4. **Batch operations are more efficient** for multiple records
5. **Always include required fields**: voyage_id, location, timestamp
6. **Weather data is optional** but useful for analysis

---

## Authentication Header

All POST requests (except /health) require authentication:

```
X-API-Key: sailing-api-key-12345
```

Or with JWT:

```
Authorization: Bearer <your-jwt-token>
```
