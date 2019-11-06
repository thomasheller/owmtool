package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getJSON(url string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("Failed to create HTTP request: %v", err)
		return err
	}

	client := http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to do HTTP request: %v", err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read body from HTTP request: %v", err)
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		log.Printf("Failed to unmarshal JSON from HTTP request: %v", err)
		return err
	}

	return nil
}
