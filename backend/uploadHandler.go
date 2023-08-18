package main

import (
	"database/sql"
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
	name := r.URL.Query().Get("name")
	description := r.URL.Query().Get("description")
	tags := r.URL.Query().Get("tags")
	type params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Tags        string `json:"tags"`
	}
	if name == "" {
		SendErr(w, http.StatusBadRequest, "Check Msg Properly")
		return
	}
	wordDir := os.Getenv("worksUploadDir")
	file, fileHandle, err := r.FormFile("art")
	if err != nil {
		SendErr(w, http.StatusBadRequest, "File Not found")
		return
	}
	image, fileHandleImg, err := r.FormFile("image")
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Image Not found")
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
	imgId := uuid.NewString()
	filePath := fmt.Sprintf("%s/%s%s", wordDir, fileId, filepath.Ext(fileHandle.Filename))
	imgPath := fmt.Sprintf("%s/%s%s", wordDir, imgId, filepath.Ext(fileHandleImg.Filename))
	diskFile, err := os.Create(filePath)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to upload Now")
		return
	}
	imgFile, err := os.Create(imgPath)
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
	_, err = io.Copy(imgFile, image)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "Unable to upload file")
		return
	}

	_, err = db.queries.UploadWork(r.Context(), database.UploadWorkParams{
		ID:         fileId,
		Userid:     userId,
		Name:       name,
		Uploadedat: time.Now().UTC(),
		Path:       filePath,
		Image:      imgPath,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		Tags: tags,
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
