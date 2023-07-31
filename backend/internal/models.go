// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"
)

type Token struct {
	Token      string
	Userid     string
	Expirydate time.Time
}

type User struct {
	ID           string
	Email        string
	Password     string
	Registeredat sql.NullTime
	Verfied      bool
}

type Work struct {
	ID         string
	Userid     string
	Name       string
	Path       string
	Uploadedat time.Time
}
