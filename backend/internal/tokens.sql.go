// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: tokens.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createToken = `-- name: CreateToken :execresult
INSERT INTO tokens (
  token,userId,expiryDate
) VALUES (
  ?,?,?
)
`

type CreateTokenParams struct {
	Token      string
	Userid     string
	Expirydate time.Time
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createToken, arg.Token, arg.Userid, arg.Expirydate)
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM tokens
WHERE userId = ?
`

func (q *Queries) DeleteToken(ctx context.Context, userid string) error {
	_, err := q.db.ExecContext(ctx, deleteToken, userid)
	return err
}

const getToken = `-- name: GetToken :one
SELECT token, userid, expirydate FROM tokens
WHERE token = ? LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, token string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getToken, token)
	var i Token
	err := row.Scan(&i.Token, &i.Userid, &i.Expirydate)
	return i, err
}

const getTokenViaEmail = `-- name: GetTokenViaEmail :one
SELECT token, userid, expirydate FROM tokens
WHERE userId = (SELECT id FROM users WHERE email = ?) LIMIT 1
`

func (q *Queries) GetTokenViaEmail(ctx context.Context, email string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getTokenViaEmail, email)
	var i Token
	err := row.Scan(&i.Token, &i.Userid, &i.Expirydate)
	return i, err
}
