package main

import (
	"errors"
	"fmt"

	"github.com/Qayyax/pokedexcli/internal"
)

func commandMap(cfg *internal.Conf) error {
	currentData, err := internal.GetData(cfg.Next)
	if err != nil {
		return err
	}
	// currentApi = currentData.Next
	cfg.Next = currentData.Next
	cfg.Previous = currentData.Previous
	for _, location := range currentData.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *internal.Conf) error {
	if cfg.Previous == nil {
		return errors.New("There is no previous page")
	}
	currentData, err := internal.GetData(*cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = currentData.Next
	cfg.Previous = currentData.Previous

	for _, location := range currentData.Results {
		fmt.Println(location.Name)
	}
	return nil

}
