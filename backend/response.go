package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
}
type ErrorResponse struct {
	Response
	Message string `json:"message"`
}
type SuccessResponse struct {
	Response
	Data interface{} `json:"data"`
}

func SendErr(w http.ResponseWriter, status int, payload string) {
	_, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Unable to parse")
		return
	}
	SendJSON(w, status, ErrorResponse{
		Message: payload,
		Response: Response{
			Success: false,
		},
	})
}
func SendSuccess(w http.ResponseWriter, status int, payload interface{}) {
	_, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Unable to parse")
		return
	}
	SendJSON(w, status, SuccessResponse{
		Response: Response{
			Success: true,
		},
		Data: payload,
	})
}

func SendJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Unable to parse")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
