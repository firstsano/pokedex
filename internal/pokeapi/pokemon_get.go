package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := pokemonUrl + "/" + name

	if dat, ok := c.cache.Get(url); ok {
		pokemon := RespPokemon{}
		if err := json.Unmarshal(dat, &pokemon); err != nil {
			return RespPokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)

	pokemon := RespPokemon{}
	if err = json.Unmarshal(dat, &pokemon); err != nil {
		return RespPokemon{}, err
	}

	return pokemon, nil
}
