package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

func (db *DbAPi) SignInHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		SendErr(w, 400, "Invalid Email")
		return
	}
	user, err := db.queries.GetUser(r.Context(), params.Email)

	if err != nil {
		SendErr(w, 400, "Not Registered")
		return
	}
	passwordCorrect := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if passwordCorrect == nil {
		token, err := db.createTokens(r.Context(), params.Email)
		if err == nil {
			SendSuccess(w, 200, token)
		}
	}

}
