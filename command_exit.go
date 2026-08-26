package main

import (
	"fmt"
	"os"

	"github.com/argwilson/pokedex/internal/pokeapi"
)

func commandExit(cfg *config, cache pokeapi.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
