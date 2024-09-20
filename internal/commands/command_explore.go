package commands

import "fmt"

func commandExplore(c *config, args []string) error {
	if len(args) > 0 {
		fmt.Printf("exploring: %v\n", args[0])
	}

	return nil
}
