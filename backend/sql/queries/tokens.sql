-- name: GetToken :one
SELECT * FROM tokens
WHERE userId = (SELECT id FROM users WHERE email = ?) LIMIT 1;

-- name: CreateToken :execresult
INSERT INTO tokens (
  token,userId,expiryDate
) VALUES (
  ?,?,?
);

-- name: DeleteToken :exec
DELETE FROM tokens
WHERE userId = ?;