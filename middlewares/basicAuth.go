package middlewares

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/Imtiaz246/Book-Server/database"
	"github.com/Imtiaz246/Book-Server/utils"
	"net/http"
	"strings"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := r.Header.Get("Authorization")
		if ai == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.CreateErrorJson(errors.New("authentication required")))
			return
		}
		// Get the token from the header.
		eAuthToken := strings.Split(r.Header.Get("Authorization"), " ")
		uAuthInfo, err := base64.StdEncoding.DecodeString(eAuthToken[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(utils.CreateErrorJson(err))
			return
		}
		cx := bytes.Index(uAuthInfo, []byte(":"))
		username := string(uAuthInfo[:cx])
		password := string(uAuthInfo[cx+1:])

		db := database.GetDB()
		db.Lock()

		if err = db.Authenticate(username, password); err != nil {
			w.Write(utils.CreateErrorJson(err))
			db.UnLock()
			return
		}
		db.UnLock()
		next.ServeHTTP(w, r)
	})
}
