package commands

import (
	"os"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandExit(api *pokeapi.PokeAPI, args []string) error {
	os.Exit(0)
	return nil
}
