package main

import (
	"os"

	"github.com/Qayyax/pokedexcli/internal"
)

func commandExit(*internal.Conf) error {
	os.Exit(0)
	return nil
}
