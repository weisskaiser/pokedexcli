package commands

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("required argument area name is missing")
	}

	fmt.Printf("Exploring %s...\n", args[0])

	root, err := c.PokemonApiClient.GetLocationPokemons(args[0])
	if err != nil {
		return err
	}

	if len(root.PokemonEncounters) == 0 {
		fmt.Println("There are no pokemons in the area")
		return nil
	}

	fmt.Println("Found Pokemon:")

	for _, e := range root.PokemonEncounters {
		fmt.Println(e.Pokemon.Name)
	}

	return nil
}
