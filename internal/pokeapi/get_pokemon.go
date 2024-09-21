package pokeapi

import (
	"encoding/json"

	httpclient "github.com/mateus-f-torres/boot_pokedex/internal/http-client"
)

type payload struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (api *PokeAPI) GetPokemon(name string) (payload, error) {
	p := payload{}

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	data, hit := api.cache.Get(url)
	if !hit {
		body, err := httpclient.Get(url)
		if err != nil {
			return p, err
		}

		data = body
		api.cache.Add(url, body)
	}

	err := json.Unmarshal(data, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}
