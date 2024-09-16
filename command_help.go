package main

import "fmt"

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	cmds := getCommands()

	for _, c := range cmds {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}

	return nil
}
