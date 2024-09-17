package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Locations struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func GetLocations(url string) (Locations, error) {
	locs := Locations{}

	res, err := http.Get(url)
	if err != nil {
		return locs, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return locs, err
	}

	err = json.Unmarshal(body, &locs)
	if err != nil {
		return locs, err
	}

	return locs, nil
}
