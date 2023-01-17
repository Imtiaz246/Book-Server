package Middleware

import (
	"net/http"
)

// AddHeaders adds some common header in the response.
func AddHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add {content-type : "application/json"}
		w.Header().Add("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
