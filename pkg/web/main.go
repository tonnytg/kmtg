package web

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/tonnytg/kmtg/domain/pod"
	"io/ioutil"
	"net/http"
	"os"
)

func Start(url string) {

	token := os.Getenv("TOKEN")

	h := http.Header{}
	h.Add("Content-Type", "application/json")
	h.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	caCert := "/tmp/ca.crt"

	// Load the CA cert
	caCertPool := x509.NewCertPool()
	caCertData, _ := ioutil.ReadFile(caCert)
	caCertPool.AppendCertsFromPEM(caCertData)

	// Create the client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	// Create the request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	_, err = prettyjson.Format(body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(s))

	var Pods pod.Pods
	json.Unmarshal(body, &Pods)
	for i, v := range Pods.Items {
		fmt.Println(i, v.Metadata.Name)
	}
}
