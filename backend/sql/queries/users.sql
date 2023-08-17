-- name: GetUser :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: GetUserViaId :many
SELECT a.id as userId,a.email as email,w.id as fileId,w.name as Name,w.uploadedAt as uploadedAt FROM users a LEFT JOIN works w ON a.id = w.userId WHERE a.id = ?;

-- name: CreateUser :execresult
INSERT INTO users (
  id,email,password
) VALUES (
  ?,?,?
);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;