-- name: GetWork :one
SELECT * FROM works
WHERE id = ? LIMIT 1;

-- name: UploadWork :execresult
INSERT INTO works (
  id,userId,name,description,tags,path,uploadedAt
) VALUES (
  ?,?,?,?,?,?,?
);

-- name: DeleteWork :exec
DELETE FROM works
WHERE id = ?;