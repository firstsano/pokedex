package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *configuration, args ...string) error {
	if len(args) != 1 {
		return errors.New("location required to explore")
	}

	location := args[0]
	locationWithPokemons, err := cfg.pokeClient.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationWithPokemons.PokemonEncounters {
		fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
