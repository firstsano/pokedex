package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *configuration) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocation)
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

func commandMapb(cfg *configuration) error {
	if cfg.prevLocation == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocation)
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
