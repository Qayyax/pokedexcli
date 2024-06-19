package main

import "fmt"

var currentApi string = "https://pokeapi.co/api/v2/location"

func commandMap() error {
	currentData, err := getData(currentApi)
	if err != nil {
		return err
	}
	currentApi = currentData.Next
	for _, location := range currentData.Results {
		fmt.Println(location.Name)
	}
	return nil
}
