CREATE TABLE event (
                               id BIGINT AUTO_INCREMENT PRIMARY KEY,
                               aggregate_id VARCHAR(255) NOT NULL DEFAULT '',
                               event_type VARCHAR(100) NOT NULL DEFAULT '',
                               payload TEXT NOT NULL DEFAULT '',
                               occurred_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               published BOOLEAN NOT NULL DEFAULT FALSE,
                               published_at TIMESTAMP NULL,
                               INDEX idx_aggregate_id (aggregate_id),
                               INDEX idx_event_type (event_type),
                               INDEX idx_published (published),
                               PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
