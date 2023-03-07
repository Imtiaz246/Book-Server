package controllers

import (
	"github.com/Imtiaz246/Book-Server/utils"
	"net/http"
)

func Ping(w http.ResponseWriter, _ *http.Request) {
	jsonData, err := utils.CreateSuccessJson("Pong")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.CreateErrorJson(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
