package commands

import "fmt"

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	cmds := GetCommands()

	for _, c := range cmds {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}

	return nil
}
