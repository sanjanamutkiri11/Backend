CREATE TABLE IF NOT EXISTS scaling_events (
    id SERIAL PRIMARY KEY,
    action VARCHAR(20) NOT NULL,
    reason TEXT,
    system_state_id INT REFERENCES system_state(id),
    triggered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
