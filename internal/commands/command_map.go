package commands

import (
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandMap(c *config) error {
	locations, err := pokeapi.GetLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
