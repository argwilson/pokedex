package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	if pokemon, ok := cfg.pokedex[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Print("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Print("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf(" -%s\n", t.Type.Name)
		}
		return nil
	}
	fmt.Printf("You have not caught that pokemon\n")

	return nil
}
