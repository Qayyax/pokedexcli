package main

func main() {

	var cfg *conf
	cfg = &conf{
		next:     "https://pokeapi.co/api/v2/location",
		previous: nil,
	}

	startRepl(cfg)
}
