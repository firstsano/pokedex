package main

import (
	"time"

	"github.com/firstsano/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	startREPL(cfg)
}
