package main

import (
	"net/http"
	"time"
)

type Work_M struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	UploadedAt time.Time `json:"uploadedAt"`
	Path       string    `json:"path"`
	Image      string    `json:"image"`
}
type User_M struct {
	Email  string   `json:"email"`
	UserId string   `json:"userId"`
	Works  []Work_M `json:"works"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("user")
	if userId == "" {
		SendErr(w, http.StatusForbidden, "Not Signed")
		return
	}
	profile, err := db.queries.GetUserViaId(r.Context(), userId)
	if err != nil {
		SendErr(w, http.StatusForbidden, "Not Signed")
		return
	}

	resData := User_M{
		Email:  profile[0].Email,
		UserId: userId,
		Works:  []Work_M{},
	}
	for _, row := range profile {
		if row.Name.Valid && row.Fileid.Valid {
			resData.Works = append(resData.Works, Work_M{
				Name:       row.Name.String,
				Id:         row.Fileid.String,
				UploadedAt: row.Uploadedat.Time,
				Path:       row.Path.String,
				Image:      row.Image.String,
			})
		}
	}
	SendSuccess(w, 200, resData)
}
