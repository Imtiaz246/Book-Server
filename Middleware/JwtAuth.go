package Middleware

import (
	"BookServer/Database"
	"BookServer/Utils"
	"bytes"
	"encoding/base64"
	"errors"
	"net/http"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := r.Header.Get("Authorization")
		if ai == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			w.Write(Utils.CreateErrorJson(errors.New("authentication required")))
			return
		}

		eAuthToken := r.Header.Get("Authorization")[7:]
		uAuthInfo, err := base64.StdEncoding.DecodeString(eAuthToken)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		cx := bytes.Index(uAuthInfo, []byte(":"))
		username := string(uAuthInfo[:cx])
		password := string(uAuthInfo[cx+1:])

		db := Database.GetDB()
		db.Lock()

		if err = db.Authenticate(username, password); err != nil {
			w.Write(Utils.CreateErrorJson(err))
			db.UnLock()
			return
		}
		db.UnLock()
		next.ServeHTTP(w, r)
	})
}
