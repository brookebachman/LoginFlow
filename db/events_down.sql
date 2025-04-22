-- Drop the 'events' table
DROP TABLE IF EXISTS events;

-- Drop indexes if needed (optional, depending on your needs)
DROP INDEX IF EXISTS idx_events_tenant_origin;
DROP INDEX IF EXISTS idx_events_timestamp;
