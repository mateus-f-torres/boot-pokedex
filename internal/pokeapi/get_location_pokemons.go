package pokeapi

import (
	"encoding/json"

	httpclient "github.com/mateus-f-torres/boot_pokedex/internal/http-client"
)

const location_url string = "https://pokeapi.co/api/v2/location"

type locationDetails struct {
	// ID    int    `json:"id"`
	// Name  string `json:"name"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

func (api *PokeAPI) GetLocationPokemons(locName string) ([]string, error) {
	pokemons := []string{}

	res, err := httpclient.Get("https://pokeapi.co/api/v2/location/" + locName)
	if err != nil {
		return pokemons, err
	}

	var locDetails locationDetails
	err = json.Unmarshal(res, &locDetails)
	if err != nil {
		return pokemons, err
	}

	// NOTE: already breaking up a little to make it easier to use goroutines
	locPokemons := map[string]bool{}
	for _, item := range locDetails.Areas {
		api.searchAreaForPokemons(item.URL, locPokemons)
	}

	for k := range locPokemons {
		pokemons = append(pokemons, k)
	}

	return pokemons, nil
}

type areaDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			// URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (api *PokeAPI) searchAreaForPokemons(rawUrl string, locPokemons map[string]bool) error {
	var areaDetails areaDetails
	rawArea, err := httpclient.Get(rawUrl)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(rawArea, &areaDetails)
	if err != nil {
		return nil
	}

	for _, encounter := range areaDetails.PokemonEncounters {
		name := encounter.Pokemon.Name
		if _, ok := locPokemons[name]; !ok {
			locPokemons[name] = true
		}
	}

	return nil
}
