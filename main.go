package main

import (
	"bytes"
	"net/http"
	"net/url"
)

func main() {

	postData := []byte("abc")

	proxyUrl, _ := url.Parse("http://localhost:3128")

	proxyValue := http.ProxyURL(proxyUrl)

	myClient := &http.Client{Transport: &http.Transport{Proxy: proxyValue}}

	req, _ := http.NewRequest("POST", "http://www.sample.com", bytes.NewReader(postData))

	req.Header.Add("Content-Type", "text/xml")

	resp, err := myClient.Do(req)

	if err != nil {

		panic(err)

		return
	}

	defer resp.Body.Close()

}
