package main

import (
	"time"

	"github.com/firstsano/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &configuration{
		pokeClient: pokeClient,
	}

	startREPL(cfg)
}
