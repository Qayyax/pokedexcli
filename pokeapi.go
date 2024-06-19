package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Locations struct {
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func getData(url string) (Locations, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Unable to process request with error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		locationData := Locations{}
		json.NewDecoder(response.Body).Decode(&locationData)
		return locationData, nil
	}
	return Locations{}, errors.New("Unable to fetch data")
}
