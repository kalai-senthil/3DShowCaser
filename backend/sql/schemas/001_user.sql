-- +goose Up

CREATE TABLE users(id VARCHAR(36) PRIMARY KEY,email VARCHAR(320) UNIQUE NOT NULL,password TEXT NOT NULL,registeredAt DATETIME DEFAULT CURRENT_TIMESTAMP,verfied BOOLEAN NOT NULL DEFAULT false);

-- +goose Down
DROP TABLE users;