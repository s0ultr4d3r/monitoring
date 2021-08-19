package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// https://api.vk.com/method/messages.send?user_id=109486533&access_token=56936fc786c413d3b5f7732a4d346b4d2e603bfb007ecfcf7e456bdeaa38fd488c3df27879ac398724322&message=test&v=5.45

func sendMsg(msg string) {

	apiStr := "https://api.vk.com/method/messages.send"
	vers := "5.45"
	userID := "109486533"
	token := "56936fc786c413d3b5f7732a4d346b4d2e603bfb007ecfcf7e456bdeaa38fd488c3df27879ac398724322"
	url := apiStr + token + vers + userID + msg

	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
}
