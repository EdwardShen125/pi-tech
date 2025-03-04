CREATE TABLE users (
                       id BIGINT NOT NULL AUTO_INCREMENT,
                       email VARCHAR(255) NOT NULL,
                       hashed_password VARCHAR(255) NOT NULL,
                       status VARCHAR(50) NOT NULL,  -- active、locked、pending 等
                       registered_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       last_login_at TIMESTAMP NULL,
                       PRIMARY KEY (id),
                       UNIQUE KEY idx_email (email)
)
    PARTITION BY HASH(id) PARTITIONS 16;
