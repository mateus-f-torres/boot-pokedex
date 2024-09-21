package commands

import (
	"fmt"
	"math/rand"

	"github.com/mateus-f-torres/boot_pokedex/internal/pokeapi"
)

const FAIL_MODIFIER float64 = 0.1

func commandCatch(api *pokeapi.PokeAPI, args []string) error {
	if len(args) < 1 {
		fmt.Println("missing pokemon")
		return nil
	}

	p, err := api.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", p.Name)
	if x := rand.Intn(p.BaseExperience); float64(x) < float64(p.BaseExperience)*FAIL_MODIFIER {
		fmt.Printf("%v escaped!\n", p.Name)
		return nil
	}

	fmt.Printf("%v was caught!\n", p.Name)
	pokemon := pokeapi.Pokemon{
		Name:   p.Name,
		Height: p.Height,
		Weight: p.Weight,
		Stats:  map[string]int{},
		Types:  []string{},
	}
	for _, stat := range p.Stats {
		pokemon.Stats[stat.Stat.Name] = stat.BaseStat
	}
	for _, t := range p.Types {
		pokemon.Types = append(pokemon.Types, t.Type.Name)
	}

	api.Pokedex[pokemon.Name] = pokemon

	return nil
}
