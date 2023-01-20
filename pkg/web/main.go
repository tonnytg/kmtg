package web

import (
	"fmt"
	"github.com/tonnytg/webreq"
	"log"
)

func Start(url string) {

	headers := webreq.NewHeaders()
	headers.Add("Content-Type", "application/json")

	request := webreq.Builder("GET")
	request.SetURL(url)

	body, err := request.Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}