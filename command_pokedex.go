package main

import "fmt"

func commandPokedex(cfg *configuration, _ ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemons yet...")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemon, _ := range cfg.caughtPokemon {
		fmt.Printf("\t- %s\n", pokemon)
	}
	return nil
}
