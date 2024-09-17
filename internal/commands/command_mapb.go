package commands

import (
	"errors"
	"fmt"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandMapB(c *config) error {
	if c.Previous == "" {
		return errors.New("no previous location list")
	}

	locations, err := pokeapi.GetLocations(c.Previous)
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
