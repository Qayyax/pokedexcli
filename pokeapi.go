package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Locations struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type conf struct {
	next     string
	previous *string
}

func getData(url string) (Locations, error) {
	response, err := http.Get(url)
	if err != nil {
		errValue := fmt.Sprintf("Unable to process request with error: %v", err)
		return Locations{}, errors.New(errValue)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		locationData := Locations{}
		json.NewDecoder(response.Body).Decode(&locationData)
		return locationData, nil
	}
	return Locations{}, errors.New("Unable to fetch data")
}
