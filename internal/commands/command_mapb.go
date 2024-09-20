package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	httpclient "github.com/mateus-f-torres/boot_pokedex/internal/http-client"
	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandMapB(c *config) error {
	if c.Previous == "" {
		return errors.New("no previous location list")
	}

	locations := pokeapi.Locations{}

	data, hit := c.Cache.Get(c.Previous)
	if !hit {
		body, err := httpclient.Get(c.Previous)
		if err != nil {
			return err
		}
		data = body
		c.Cache.Add(c.Previous, body)
	}

	err := json.Unmarshal(data, &locations)
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
