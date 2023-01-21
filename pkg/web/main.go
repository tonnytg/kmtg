package web

import (
	"fmt"
	"net/http"
)

func Start(url string, token string) {

	h := http.Header{}
	h.Add("Content-Type", "application/json")
	h.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	body, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)

}
