package main

import "fmt"

// var currentApi string = "https://pokeapi.co/api/v2/location"

func commandMap(cfg *conf) error {
	currentData, err := getData(cfg.next)
	if err != nil {
		return err
	}
	// currentApi = currentData.Next
	cfg.next = currentData.Next
	cfg.previous = currentData.Previous
	for _, location := range currentData.Results {
		fmt.Println(location.Name)
	}
	return nil
}
