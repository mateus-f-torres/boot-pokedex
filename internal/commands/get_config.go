package commands

import "github.com/mateus-f-torres/boot_pokedex/internal/cache"

type config struct {
	Next     string
	Previous string
	Cache    cache.Cache
}

func GetConfig(cache cache.Cache) config {
	return config{
		Next:     "https://pokeapi.co/api/v2/location",
		Previous: "",
		Cache:    cache,
	}
}
