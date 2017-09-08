package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"number": 1}`)
}

func TestGetAstronauts(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handler))
	hc := server.Client()

	client := LiveGetWebRequest{
		client: *hc,
	}
	client.url = server.URL

	amount := GetAstronauts(client)
	if amount != 1 {
		t.Errorf("People in space, got: %d, want: %d.", amount, 1)
	}
}
