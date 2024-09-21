package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandHelp(api *pokeapi.PokeAPI, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	cmds := GetCommands()

	pad := 0
	slc := []Command{}
	for _, v := range cmds {
		if len(v.name) > pad {
			pad = len(v.name)
		}

		slc = append(slc, v)
	}

	slices.SortFunc(slc, func(a, b Command) int {
		return a.order - b.order
	})

	for _, v := range slc {
		name := v.name
		if len(name) < pad {
			name = name + strings.Repeat(" ", pad-len(name))
		}
		fmt.Printf("  %v   %v\n", name, v.description)
	}

	return nil
}
