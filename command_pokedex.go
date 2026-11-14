package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Print("Your Pokedex:\n")
	for pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
