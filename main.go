package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand) error
}

func main() {
	cliKeys := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for reader.Scan() {
		userInput := reader.Text()
		if command, ok := cliKeys[userInput]; ok {
			command.callback(cliKeys)
			fmt.Print("Pokedex > ")
		} else {
			fmt.Println("Incorrect command")
			fmt.Print("Pokedex > ")
		}
	}

}

func commandHelp(command map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	fmt.Println("help:", command["help"].description)
	fmt.Println("exit:", command["exit"].description)
	fmt.Println()
	return nil
}

func commandExit(command map[string]cliCommand) error {
	os.Exit(0)
	return nil
}
