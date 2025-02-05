package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemonapiLocationsRoot struct {
	Next     string               `json:"next"`
	Previous string               `json:"previous"`
	Results  []pokemonapiLocation `json:"results"`
}

type pokemonapiLocation struct {
	Name string `json:"name"`
}

type pokemonapiPokemonRoot struct {
	Name           string  `json:"name"`
	BaseExperience float64 `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Types          []pokemonTypeWrapper
	Stats          []pokemonStatWrapper
}

type pokemonTypeWrapper struct {
	Type NameEntity `json:"type"`
}

type pokemonStatWrapper struct {
	Stat     NameEntity `json:"stat"`
	BaseStat int        `json:"base_stat"`
}

type pokemonapiLocationRoot struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon NameEntity `json:"pokemon"`
}

type NameEntity struct {
	Name string `json:"name"`
}

type PokemonApiClient struct {
	cache   *Cache
	baseUrl string
}

func NewPokemonApiClient(cache *Cache, baseUrl string) *PokemonApiClient {
	return &PokemonApiClient{cache, baseUrl}
}

func (client *PokemonApiClient) GetLocations(locationsUrl string) (pokemonapiLocationsRoot, error) {

	var content []byte
	if val, ok := client.cache.Get(locationsUrl); ok {
		content = val
	} else {
		res, err := http.Get(locationsUrl)
		if err != nil {
			return pokemonapiLocationsRoot{}, fmt.Errorf("failed to make network call: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return pokemonapiLocationsRoot{}, fmt.Errorf("http call failed: %s", res.Status)
		}
		content, err = io.ReadAll(res.Body)
		if err != nil {
			return pokemonapiLocationsRoot{}, fmt.Errorf("could not read response: %w", err)
		}
		client.cache.Add(locationsUrl, content)
	}
	var root pokemonapiLocationsRoot
	if err := json.Unmarshal(content, &root); err != nil {
		return pokemonapiLocationsRoot{}, fmt.Errorf("error deserializing body: %s", err)
	}

	return root, nil
}

func (client *PokemonApiClient) GetLocationPokemons(locationName string) (pokemonapiLocationRoot, error) {

	var content []byte
	if val, ok := client.cache.Get(locationName); ok {
		content = val
	} else {
		res, err := http.Get(client.baseUrl + "/api/v2/location-area/" + locationName)
		if err != nil {
			return pokemonapiLocationRoot{}, fmt.Errorf("failed to make network call: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return pokemonapiLocationRoot{}, fmt.Errorf("http call failed: %s", res.Status)
		}
		content, err = io.ReadAll(res.Body)
		if err != nil {
			return pokemonapiLocationRoot{}, fmt.Errorf("could not read response: %w", err)
		}
		client.cache.Add(locationName, content)
	}
	var root pokemonapiLocationRoot
	if err := json.Unmarshal(content, &root); err != nil {
		return pokemonapiLocationRoot{}, fmt.Errorf("error deserializing body: %s", err)
	}

	return root, nil
}

func (client *PokemonApiClient) GetPokemonInformation(pokemonName string) (pokemonapiPokemonRoot, error) {
	var content []byte
	if val, ok := client.cache.Get(pokemonName); ok {
		content = val
	} else {
		res, err := http.Get(client.baseUrl + "/api/v2/pokemon/" + pokemonName)
		if err != nil {
			return pokemonapiPokemonRoot{}, fmt.Errorf("failed to make network call: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return pokemonapiPokemonRoot{}, fmt.Errorf("http call failed: %s", res.Status)
		}
		content, err = io.ReadAll(res.Body)
		if err != nil {
			return pokemonapiPokemonRoot{}, fmt.Errorf("could not read response: %w", err)
		}
		client.cache.Add(pokemonName, content)
	}
	var root pokemonapiPokemonRoot
	if err := json.Unmarshal(content, &root); err != nil {
		return pokemonapiPokemonRoot{}, fmt.Errorf("error deserializing body: %s", err)
	}

	return root, nil
}
