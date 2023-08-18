// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  id,email,password
) VALUES (
  ?,?,?
)
`

type CreateUserParams struct {
	ID       string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.ID, arg.Email, arg.Password)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password, registeredat, verfied FROM users
WHERE email = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Registeredat,
		&i.Verfied,
	)
	return i, err
}

const getUserViaId = `-- name: GetUserViaId :many
SELECT a.id as userId,a.email as email,w.id as fileId,w.name as Name,w.uploadedAt as uploadedAt,w.path as path,w.image as image FROM users a LEFT JOIN works w ON a.id = w.userId WHERE a.id = ?
`

type GetUserViaIdRow struct {
	Userid     string
	Email      string
	Fileid     sql.NullString
	Name       sql.NullString
	Uploadedat sql.NullTime
	Path       sql.NullString
	Image      sql.NullString
}

func (q *Queries) GetUserViaId(ctx context.Context, id string) ([]GetUserViaIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserViaId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserViaIdRow
	for rows.Next() {
		var i GetUserViaIdRow
		if err := rows.Scan(
			&i.Userid,
			&i.Email,
			&i.Fileid,
			&i.Name,
			&i.Uploadedat,
			&i.Path,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
