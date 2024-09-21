package commands

import (
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandInspect(api *pokeapi.PokeAPI, args []string) error {
	if len(args) < 1 {
		fmt.Println("missing pokemon")
		return nil
	}

	p, ok := api.Pokedex[args[0]]
	if !ok {
		fmt.Printf("you have not caught %v yet\n", args[0])
		return nil
	}

	fmt.Printf(`Name: %v
Height: %v
Weight: %v
Stats:
  -hp: %v
  -attack: %v
  -defense: %v
  -special-attack: %v
  -special-defense: %v
  -speed: %v
`,
		p.Name,
		p.Height,
		p.Weight,
		p.Stats["hp"],
		p.Stats["attack"],
		p.Stats["defense"],
		p.Stats["special-attack"],
		p.Stats["special-defense"],
		p.Stats["speed"],
	)

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %v\n", t)
	}

	return nil
}
