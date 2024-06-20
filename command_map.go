package main

import (
	"errors"
	"fmt"
)

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

func commandMapb(cfg *conf) error {
	if cfg.previous == nil {
		return errors.New("There is no previous page")
	}
	currentData, err := getData(*cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = currentData.Next
	cfg.previous = currentData.Previous

	for _, location := range currentData.Results {
		fmt.Println(location.Name)
	}
	return nil

}
