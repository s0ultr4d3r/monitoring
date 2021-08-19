package main

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// https://monitoring.mvk.com/states/api/?api_app=states_api&api_key=2e219f0f7e9bc9856dc40166505213c3fbc888b9
func reqStr() string {
	apiAddr := "https://monitoring.mvk.com/states/api/?api_app=states_api"
	apiKey := "2e219f0f7e9bc9856dc40166505213c3fbc888b9"
	reqStr := apiAddr + "&api_key=" + apiKey
	return (reqStr)
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func apiRq() *MonAnsw {

	monAnsw := new(MonAnsw)
	err := getJson(reqStr(), monAnsw)
	if err != nil {
		log.Fatal(err)
	}
	return monAnsw
}
