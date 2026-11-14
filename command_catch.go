package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	var p int
	switch {
	case pokemon.BaseExperience < 100:
		p = rand.Intn(100)
	default:
		p = rand.Intn(pokemon.BaseExperience)
	}
	if p < 50 {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught\n", pokemon.Name)
		fmt.Print("You may now inspect it with the inspect command.\n")
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}
