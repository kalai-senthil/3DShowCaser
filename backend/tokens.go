package main

import (
	"context"
	"fmt"
	"os"
	"time"

	database "github.com/3D-ShowCaser/backend/internal"
	"github.com/golang-jwt/jwt"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func (db *DbAPi) createTokens(Context context.Context, email string) (Token, error) {
	tokensFromDB, err := db.queries.GetToken(Context, email)
	if err == nil {
		return Token{AccessToken: tokensFromDB.Token}, nil
	}
	userFromDB, err := db.queries.GetUser(Context, email)
	if err != nil {
		return Token{}, fmt.Errorf("%s not registered", userFromDB.Email)
	}
	expiryDate := time.Now().Add(24 * time.Hour * 7)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expiryDate

	access_token, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return Token{}, err
	}

	_, err = db.queries.CreateToken(Context, database.CreateTokenParams{
		Token:      access_token,
		Userid:     userFromDB.ID,
		Expirydate: expiryDate,
	})
	if err != nil {
		return Token{}, err
	}
	return Token{AccessToken: access_token}, nil
}
