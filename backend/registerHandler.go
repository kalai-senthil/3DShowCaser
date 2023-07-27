package main

import (
	"encoding/json"
	"net/http"

	database "github.com/3D-ShowCaser/backend/internal"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (db *DbAPi) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Id       string `id`
		Email    string `email`
		Password string `password`
	}
	decoder := json.NewDecoder(r.Body)
	params := User{}
	decoder.Decode(&params)
	if params.Email == "" {
		SendErr(w, 400, "Email is invalid")
		return
	}
	if params.Password == "" {
		SendErr(w, 400, "Password is invalid")
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		params.Password = string(hashedPass)
	}
	id := uuid.New()
	res, err := db.queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:       []byte(id.String()),
		Email:    params.Email,
		Password: string(hashedPass),
	})
	if err != nil {
		SendErr(w, 400, "Registration Unsuccessful")
		return
	}
	SendSuccess(w, 201, res)
}
