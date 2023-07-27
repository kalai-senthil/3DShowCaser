package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	database "github.com/3D-ShowCaser/backend/internal"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type DbAPi struct {
	queries *database.Queries
}

func main() {
	godotenv.Load()
	queries, err := ConnectToDB()
	if err != nil {
		log.Fatal("Unable to connect to database")
	}
	db := DbAPi{queries: queries}
	router := mux.NewRouter()
	router.Handle("/auth", GetAuthRouter(router, &db))
	router.HandleFunc("/", homeHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
