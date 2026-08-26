package main

import (
	"errors"
	"fmt"

	"github.com/argwilson/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config, cache pokeapi.Cache) error {

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, cache)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, cache pokeapi.Cache) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL, cache)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
