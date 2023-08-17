package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAuthRouter(router *mux.Router, db *DbAPi) *mux.Router {
	AuthRouter := router.PathPrefix("/auth").Subrouter()
	AuthRouter.HandleFunc("/signIn", db.SignInHandler).Methods(http.MethodPost)
	AuthRouter.HandleFunc("/register", db.RegisterHandler).Methods(http.MethodPost)
	return AuthRouter
}
