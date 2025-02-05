package commands

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("required argument pokemon name is missing")
	}

	pokemonName := args[0]

	if pokemon, ok := c.PokemonsCaught[pokemonName]; ok {

		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", s.Name, s.Value)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("-%s\n", t)
		}
		return nil
	}

	fmt.Println("you have not caught that pokemon")

	return nil
}
