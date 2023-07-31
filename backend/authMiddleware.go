package main

import (
	"fmt"
	"net/http"
	"os"
)

func (dBApi *DbAPi) authenticatedMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		if Is_URL_Public(r.RequestURI) {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Origin", os.Getenv("FRONTEND_URL"))
		token := r.Header.Get("Authorization")
		if token == "" {
			cookie, err := r.Cookie("access_token")
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			token = cookie.Value
		}
		tokensFromDB, err := dBApi.queries.GetToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		r.Header.Set("user", tokensFromDB.Userid)
		next.ServeHTTP(w, r)
	})
}
