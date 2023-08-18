-- +goose Up

CREATE TABLE works(id VARCHAR(36) PRIMARY KEY,userId VARCHAR(36) NOT NULL,name VARCHAR(100) NOT NULL,description TEXT ,tags VARCHAR(100) NOT NULL DEFAULT '',path VARCHAR(50) NOT NULL,image VARCHAR(50) NOT NULL,uploadedAt DATETIME NOT NULL, FOREIGN KEY works(userId) REFERENCES users(id), UNIQUE (userId,name));

-- +goose Down
DROP TABLE works;