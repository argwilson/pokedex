package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PokemonEncounterStruct struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandMap(cfg *config) error {
	var locations []string
	i := 1
	for i < 21 {
		current, err := strconv.Atoi(strings.Split(cfg.next, "/")[6])
		if err != nil {
			return err
		}
		res, err := http.Get(cfg.next)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return err
		}

		location := PokemonEncounterStruct{}
		err = json.Unmarshal(body, &location)
		if err != nil {
			return err
		}
		locations = append(locations, location.Name)
		cfg.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", current+1)
		cfg.previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", current)
		i++
	}

	for _, loc := range locations {
		fmt.Printf("%s\n", loc)
	}

	return nil
}
