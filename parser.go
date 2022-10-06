package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Parse(res *http.Response) string {
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	result, err := doc.Find("head title").Html()
	if err != nil {
		log.Fatal(err)
	}

	return result
}
