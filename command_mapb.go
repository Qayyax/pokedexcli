package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *conf) error {
	if cfg.previous == nil {
		return errors.New("There is no previous page")
	}
	currentData, err := getData(*cfg.previous)
	if err != nil {
		return err
	}

	for _, location := range currentData.Results {
		fmt.Println(location.Name)
	}
	return nil

}
