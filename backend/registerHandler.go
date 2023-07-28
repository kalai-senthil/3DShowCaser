package main

import (
	"encoding/json"
	"net/http"
	"strings"

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
	_, err = db.queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:       id.String(),
		Email:    params.Email,
		Password: string(hashedPass),
	})
	if err != nil {
		alreadyRegisterd := strings.Contains(err.Error(), "Duplicate entry")
		msg := "Registration Unsuccessful"
		if alreadyRegisterd {
			msg = "Already Registered"
		}
		SendErr(w, 400, msg)
		return
	}
	type Data struct {
		data interface{} `json:"data"`
	}
	SendSuccess(w, 201, Data{
		data: "Registered Successfully",
	})
}
