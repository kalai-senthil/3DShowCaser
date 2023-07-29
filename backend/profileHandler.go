package main

import (
	"net/http"
)

type User_M struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("user")
	if user == "" {
		SendErr(w, http.StatusForbidden, "Not Signed")
		return
	}
	profile, err := db.queries.GetUserViaId(r.Context(), user)
	if err != nil {
		SendErr(w, http.StatusForbidden, "Not Signed")
		return
	}

	SendSuccess(w, 200, User_M{
		Email:  profile.Email,
		UserId: profile.ID,
	})
}
