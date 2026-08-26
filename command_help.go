package main

import (
	"fmt"

	"github.com/argwilson/pokedex/internal/pokeapi"
)

func commandHelp(cfg *config, cache pokeapi.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
