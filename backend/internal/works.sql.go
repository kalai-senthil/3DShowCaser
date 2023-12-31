// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: works.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const deleteWork = `-- name: DeleteWork :exec
DELETE FROM works
WHERE id = ?
`

func (q *Queries) DeleteWork(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteWork, id)
	return err
}

const getWork = `-- name: GetWork :one
SELECT id, userid, name, description, tags, path, image, uploadedat FROM works
WHERE id = ? LIMIT 1
`

func (q *Queries) GetWork(ctx context.Context, id string) (Work, error) {
	row := q.db.QueryRowContext(ctx, getWork, id)
	var i Work
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Description,
		&i.Tags,
		&i.Path,
		&i.Image,
		&i.Uploadedat,
	)
	return i, err
}

const uploadWork = `-- name: UploadWork :execresult
INSERT INTO works (
  id,userId,name,description,tags,path,image,uploadedAt
) VALUES (
  ?,?,?,?,?,?,?,?
)
`

type UploadWorkParams struct {
	ID          string
	Userid      string
	Name        string
	Description sql.NullString
	Tags        string
	Path        string
	Image       string
	Uploadedat  time.Time
}

func (q *Queries) UploadWork(ctx context.Context, arg UploadWorkParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, uploadWork,
		arg.ID,
		arg.Userid,
		arg.Name,
		arg.Description,
		arg.Tags,
		arg.Path,
		arg.Image,
		arg.Uploadedat,
	)
}
