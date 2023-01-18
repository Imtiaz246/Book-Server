package Middleware

import (
	"BookServer/Utils"
	"context"
	"errors"
	"net/http"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := r.Header.Get("Authorization")
		if len(ai) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(Utils.CreateErrorJson(errors.New("authentication required")))
			return
		}

		eAuthToken := r.Header.Get("Authorization")[7:]
		username, err := Utils.CheckJWTValidation(eAuthToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(Utils.CreateErrorJson(err))
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
