package middlewares

import (
	"context"
	"errors"
	"github.com/Imtiaz246/Book-Server/utils"
	"net/http"
	"strings"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := r.Header.Get("Authorization")
		if len(ai) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.CreateErrorJson(errors.New("authentication required")))
			return
		}

		eAuthToken := strings.Split(r.Header.Get("Authorization"), " ")[1]
		username, err := utils.CheckJWTValidation(eAuthToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.CreateErrorJson(err))
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
