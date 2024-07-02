package main

import "github.com/Qayyax/pokedexcli/internal"

func main() {

	var cfg *internal.Conf
	cfg = &internal.Conf{
		Next:     "https://pokeapi.co/api/v2/location",
		Previous: nil,
	}

	startRepl(cfg)
}
