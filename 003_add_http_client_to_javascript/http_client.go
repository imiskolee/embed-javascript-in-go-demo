package main

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct {

}

func (h *HttpClient) Get(url string) string {
	resp,err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}
