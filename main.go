package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type command struct {
	name        string
	description string
	callback    func() error
}

func main() {
	prompt := "pokedex > "
	buf := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	for {
		fmt.Print(prompt)
		buf.Scan()
		input := buf.Text()
		c, ok := cmds[input]
		if ok {
			err := c.callback()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("unknown command: %v\n", input)
			fmt.Println("type 'help' to see list of available commands")
		}
	}
}
