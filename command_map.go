package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *configuration, _ ...string) error {
	locationsResp, err := cfg.pokeClient.ListLocations(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locationsResp.Next
	cfg.prevLocation = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *configuration, _ ...string) error {
	if cfg.prevLocation == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeClient.ListLocations(cfg.prevLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locationResp.Next
	cfg.prevLocation = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
