package commands

import (
	"fmt"
)

func commandHelp(c *config, args []string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range cliCommands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}

	return nil
}
