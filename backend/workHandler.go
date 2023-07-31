package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func workGetHandler(w http.ResponseWriter, r *http.Request) {
	artId := mux.Vars(r)["artId"]
	res, err := db.queries.GetWork(r.Context(), artId)
	if err != nil {
		SendErr(w, http.StatusNotFound, fmt.Sprintf("Art with %v not found", artId))
	}
	SendSuccess(w, 200, Work_M{
		Id:         res.ID,
		Name:       res.Name,
		UploadedAt: res.Uploadedat,
		Path:       res.Path,
	})
}
