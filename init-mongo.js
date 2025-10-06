db = db.getSiblingDB('sailing_db');

db.createCollection('voyages');
db.createCollection('checkpoints');
db.createCollection('gps_tracks');

// Create indexes
db.voyages.createIndex({ "voyage_id": 1 }, { unique: true });
db.voyages.createIndex({ "ship_id": 1 });
db.voyages.createIndex({ "departure_time": 1 });
db.voyages.createIndex({ "arrival_time": 1 });

db.checkpoints.createIndex({ "voyage_id": 1 });
db.checkpoints.createIndex({ "timestamp": 1 });

db.gps_tracks.createIndex({ "voyage_id": 1 });
db.gps_tracks.createIndex({ "timestamp": 1 });

print('Database initialized successfully');
