package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Fetch(url string) *http.Response {
	if !(strings.HasPrefix(url, "http://") && strings.HasPrefix(url, "https://")) {
		url = fmt.Sprintf("http://%s", url)
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	return res
}
