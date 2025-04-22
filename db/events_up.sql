-- Create the 'events' table
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tenant_id TEXT NOT NULL,
    username TEXT NOT NULL,
    login_status TEXT NOT NULL,
    origin TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tenant_id, username, timestamp) -- Ensure no duplicate events for the same user and tenant
);

-- Create an index on 'tenant_id' and 'origin' for faster queries on suspicious events
CREATE INDEX IF NOT EXISTS idx_events_tenant_origin ON events(tenant_id, origin);

-- Create an index on 'timestamp' to speed up time window queries
CREATE INDEX IF NOT EXISTS idx_events_timestamp ON events(timestamp);
