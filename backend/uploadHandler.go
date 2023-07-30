package main

import (
	"net/http"
	"strings"
	"time"

	database "github.com/3D-ShowCaser/backend/internal"
	"github.com/google/uuid"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("user")
	file, fileHandle, err := r.FormFile("file")
	if err != nil {
		SendErr(w, http.StatusBadRequest, "File Not found")
		return
	}
	defer file.Close()
	if userId == "" {
		SendErr(w, http.StatusForbidden, "Not Signed")
		return
	}
	fileBinary := make([]byte, fileHandle.Size)
	_, err = file.Read(fileBinary)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to parse file")
		return
	}
	fileId := uuid.NewString()
	_, err = db.queries.UploadWork(r.Context(), database.UploadWorkParams{
		ID:         fileId,
		Userid:     userId,
		Name:       fileHandle.Filename,
		Uploadedat: time.Now().UTC(),
		File:       fileBinary,
	})
	if err != nil {
		msg := "Uplod Error"
		if strings.Contains(err.Error(), "Duplicate entry") {
			msg = "Already Uploaded"
		}
		SendErr(w, http.StatusBadRequest, msg)
		return
	}
	SendSuccess(w, 200, struct {
		FileId  string `json:"id"`
		Message string `json:"message"`
		Name    string `json:"name"`
	}{FileId: fileId, Message: "File Uploaded Successfully", Name: fileHandle.Filename})
}
