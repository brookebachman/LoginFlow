CREATE TABLE IF NOT EXISTS logins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tenant_id TEXT,
    user_id TEXT,
    source_ip TEXT,
    status TEXT,
    timestamp DATETIME
);
