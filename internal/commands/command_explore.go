package commands

import (
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandExplore(api *pokeapi.PokeAPI, args []string) error {
	if len(args) > 0 {
		fmt.Printf("exploring %v...\n", args[0])
	}

	pokemons, err := api.GetLocationPokemons(args[0])
	if err != nil {
		return err
	}

	for i, p := range pokemons {
		fmt.Printf(" %v. %v\n", i+1, p)
	}

	return nil
}
