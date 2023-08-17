package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	database "github.com/3D-ShowCaser/backend/internal"
	"github.com/google/uuid"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	userId := r.Header.Get("user")
	type params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Tags        string `json:"tags"`
	}
	args := params{}
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Check Msg Properly")
		return
	}
	wordDir := os.Getenv("worksUploadDir")
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
	err = os.MkdirAll(wordDir, 0666)

	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to upload file")
		return
	}
	fileId := uuid.NewString()
	path := fmt.Sprintf("%s/%s%s", wordDir, fileId, filepath.Ext(fileHandle.Filename))
	diskFile, err := os.Create(path)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to upload Now")
		return
	}
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to parse file")
		return
	}
	_, err = io.Copy(diskFile, file)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to upload file")
		return
	}

	_, err = db.queries.UploadWork(r.Context(), database.UploadWorkParams{
		ID:         fileId,
		Userid:     userId,
		Name:       args.Name,
		Uploadedat: time.Now().UTC(),
		Path:       path,
		Description: sql.NullString{
			String: args.Description,
			Valid:  args.Description != "",
		},
		Tags: args.Tags,
	})
	if err != nil {
		fmt.Println(err)
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
