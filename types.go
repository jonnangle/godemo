package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type people struct {
	Number int `json:"number"`
}

type GetWebRequest interface {
	FetchBytes() []byte
}

type LiveGetWebRequest struct {
	GetWebRequest
	url    string
	client http.Client
}

func (l LiveGetWebRequest) FetchBytes() []byte {
	req, err := http.NewRequest(http.MethodGet, l.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	fmt.Println("***")
	fmt.Println(l.client.Timeout)
	res, getErr := l.client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}
