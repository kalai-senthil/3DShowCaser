package main

import (
	"database/sql"
	"os"

	DB "github.com/3D-ShowCaser/backend/internal"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*DB.Queries, error) {
	DB_URL := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", DB_URL)
	if err != nil {
		return nil, err
	}
	return DB.New(db), nil
}
