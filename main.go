package main

import (
	"time"

	"github.com/firstsano/pokedex/internal/pokeapi"
)

type configuration struct {
	caughtPokemon map[string]pokeapi.RespPokemon
	pokeClient    pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &configuration{
		caughtPokemon: make(map[string]pokeapi.RespPokemon),
		pokeClient:    pokeClient,
	}

	startREPL(cfg)
}
