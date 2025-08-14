package main

import (
	"time"

	"github.com/firstsano/pokedex/internal/pokeapi"
)

type Pokemon struct {
	id             int
	name           string
	baseExperience int
	height         int
	weight         int
}

type configuration struct {
	caughtPokemon map[string]Pokemon
	pokeClient    pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &configuration{
		caughtPokemon: make(map[string]Pokemon),
		pokeClient:    pokeClient,
	}

	startREPL(cfg)
}
