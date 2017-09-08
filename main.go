package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetAstronauts(getWebRequest LiveGetWebRequest) int {
	//	getWebRequest.url = "http://api.open-notify.org/astros.json"

	bodyBytes := getWebRequest.FetchBytes()
	fmt.Println(string(bodyBytes))
	peopleResult := people{}
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return peopleResult.Number
}

func main() {
	liveClient := LiveGetWebRequest{
		client: http.Client{
			Timeout: time.Second * 2, // Maximum of 2 secs
		},
	}
	liveClient.url = "http://api.open-notify.org/astros.json"

	number := GetAstronauts(liveClient)

	fmt.Println(number)
}
