package pokeapi

import (
	"encoding/json"

	httpclient "github.com/mateus-f-torres/boot_pokedex/internal/http-client"
)

type Locations struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (api *PokeAPI) GetLocations(url string) (Locations, error) {
	locs := Locations{}

	data, hit := api.cache.Get(url)
	if !hit {
		body, err := httpclient.Get(url)
		if err != nil {
			return locs, err
		}

		data = body
		api.cache.Add(url, body)
	}

	err := json.Unmarshal(data, &locs)
	if err != nil {
		return locs, err
	}

	api.Next = locs.Next
	api.Prev = locs.Previous

	return locs, nil
}
