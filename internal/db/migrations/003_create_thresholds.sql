CREATE TABLE IF NOT EXISTS thresholds (
    id SERIAL PRIMARY KEY,
    metric_type VARCHAR(20) NOT NULL,
    min_value NUMERIC(5,2) NOT NULL,
    avg_value NUMERIC(5,2) NOT NULL,
    max_value NUMERIC(5,2) NOT NULL,
    is_manual SMALLINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
