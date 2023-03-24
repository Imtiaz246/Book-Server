package controllers

import (
	"github.com/Imtiaz246/Book-Server/utils"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

func Ping(w http.ResponseWriter, _ *http.Request) {
	jsonData, err := utils.CreateSuccessJson("Pong")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.CreateErrorJson(err))
		utils.TotalResponsesWithStatusCode.With(prometheus.Labels{"code": "500"}).Inc()
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	utils.TotalResponsesWithStatusCode.With(prometheus.Labels{"code": "200"}).Inc()
}
