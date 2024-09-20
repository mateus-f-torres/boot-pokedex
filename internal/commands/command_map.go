package commands

import (
	"encoding/json"
	"fmt"

	httpclient "github.com/mateus-f-torres/boot_pokedex/internal/http-client"
	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

func commandMap(c *config) error {
	locations := pokeapi.Locations{}

	data, hit := c.Cache.Get(c.Next)
	if !hit {
		body, err := httpclient.Get(c.Next)
		if err != nil {
			return err
		}
		data = body
		c.Cache.Add(c.Next, body)
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
