-- +goose Up

CREATE TABLE tokens(token VARCHAR(200) PRIMARY KEY,userId VARCHAR(36) NOT NULL,expiryDate DATETIME NOT NULL, FOREIGN KEY tokens(userId) REFERENCES users(id));

-- +goose Down
DROP TABLE tokens;