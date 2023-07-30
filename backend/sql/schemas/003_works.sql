-- +goose Up

CREATE TABLE works(id VARCHAR(36) PRIMARY KEY,userId VARCHAR(36) NOT NULL,name VARCHAR(100) NOT NULL,file LONGBLOB NOT NULL,uploadedAt DATETIME NOT NULL, FOREIGN KEY works(userId) REFERENCES users(id), UNIQUE (userId,name));

-- +goose Down
DROP TABLE works;