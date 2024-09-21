package commands

import (
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandPokedex(api *pokeapi.PokeAPI, args []string) error {
	fmt.Println("Your Pokedex:")
	for k := range api.Pokedex {
		fmt.Printf(" - %v\n", k)
	}

	return nil
}
