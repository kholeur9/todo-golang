package middleware

import "net/http"

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "UnAuthorized", http.StatusInternalServerError)
			return
		}

		next(w, r)
	}
}