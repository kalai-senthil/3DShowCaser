-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  id,email,password
) VALUES (
  ?,?,?
);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;