package commands

import (
	"errors"
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandMapB(api *pokeapi.PokeAPI, args []string) error {
	if api.Prev == "" {
		return errors.New("no previous location list")
	}

	locations, err := api.GetLocations(api.Prev)
	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
