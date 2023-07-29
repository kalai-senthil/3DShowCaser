package main

import (
	"net/http"
)

func (dBApi *DbAPi) authenticatedMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Is_URL_Public(r.RequestURI) {
			next.ServeHTTP(w, r)
			return
		}
		token := r.Header.Get("access_token")
		if token == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		tokensFromDB, err := dBApi.queries.GetToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		r.Header.Set("user", tokensFromDB.Token)
		next.ServeHTTP(w, r)
	})
}
