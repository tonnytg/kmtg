package cmd

import (
	"flag"
	"github.com/tonnytg/kmtg/pkg/web"
	"log"
)

func Start() {
	url := flag.String("url", "", "url to call")
	method := flag.String("method", "GET", "method to call")

	flag.Parse()

	if *url == "" {
		panic("url is required")
	}
	if *method == "" {
		panic("method is required")
	}

	if *method == "GET" {
		log.Default().Println("Calling GET method")
		web.Start(*url)
	}
}
