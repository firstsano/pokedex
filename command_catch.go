package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *configuration, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	throwValue := rand.Intn(pokemon.BaseExperience)
	catched := throwValue > pokemon.BaseExperience/2
	if !catched {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[pokemonName] = pokemon
	return nil
}
