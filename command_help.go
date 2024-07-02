package main

import (
	"fmt"

	"github.com/Qayyax/pokedexcli/internal"
)

func commandHelp(conf *internal.Conf) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
