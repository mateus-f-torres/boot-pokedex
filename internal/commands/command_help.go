package commands

import (
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandHelp(api *pokeapi.PokeAPI, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	cmds := GetCommands()

	for _, c := range cmds {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}

	return nil
}
