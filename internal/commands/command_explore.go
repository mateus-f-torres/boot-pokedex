package commands

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/mateus-f-torres/boot_pokedex/internal/http-client"
)

const LOCATION_URL string = "https://pokeapi.co/api/v2/location"

type LocationDetails struct {
	// ID    int    `json:"id"`
	// Name  string `json:"name"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

type AreaDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			// URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, args []string) error {
	if len(args) > 0 {
		fmt.Printf("exploring %v...\n", args[0])
	}

	res, err := httpclient.Get(LOCATION_URL + "/" + args[0])
	if err != nil {
		return err
	}

	var locDetails LocationDetails
	err = json.Unmarshal(res, &locDetails)
	if err != nil {
		return err
	}

	pokemons := map[string]bool{}
	for _, item := range locDetails.Areas {
		var areaDetails AreaDetails
		rawArea, err := httpclient.Get(item.URL)
		if err != nil {
			return nil
		}

		err = json.Unmarshal(rawArea, &areaDetails)
		if err != nil {
			return nil
		}

		for _, encounter := range areaDetails.PokemonEncounters {
			name := encounter.Pokemon.Name
			if _, ok := pokemons[name]; !ok {
				pokemons[name] = true
			}
		}
	}

	for k := range pokemons {
		fmt.Printf(" - %v\n", k)
	}

	return nil
}
