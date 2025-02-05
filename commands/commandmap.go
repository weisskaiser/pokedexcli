package commands

import (
	"fmt"
)

func commandMap(c *config, args []string) error {

	if c.Next == "" {
		fmt.Println("Nothing else to show")
		return nil
	}

	root, err := c.PokemonApiClient.GetLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = root.Next
	c.Previous = root.Previous

	for _, v := range root.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(c *config, args []string) error {

	if c.Previous == "" {
		fmt.Println("Nothing to show")
		return nil
	}

	root, err := c.PokemonApiClient.GetLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = root.Next
	c.Previous = root.Previous

	for _, v := range root.Results {
		fmt.Println(v.Name)
	}

	return nil
}
