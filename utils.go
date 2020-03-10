package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func http_get(url string) (Response){

	var response Response

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &response)

	return response
}