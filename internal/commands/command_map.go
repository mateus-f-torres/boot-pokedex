package commands

import (
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandMap(api *pokeapi.PokeAPI, args []string) error {
	locations, err := api.GetLocations(api.Next)
	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
