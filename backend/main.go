package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
