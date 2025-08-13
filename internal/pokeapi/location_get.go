package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(location string) (RespLocationWithPokemons, error) {
	url := locationAreaURL + "/" + location
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationWithPokemons{}, err
	}

	if val, ok := c.cache.Get(url); ok {
		locationsWithPokemons := RespLocationWithPokemons{}
		if err := json.Unmarshal(val, &locationsWithPokemons); err != nil {
			return RespLocationWithPokemons{}, err
		}

		return locationsWithPokemons, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationWithPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationWithPokemons{}, err
	}

	locationsWithPokemons := RespLocationWithPokemons{}
	if err = json.Unmarshal(dat, &locationsWithPokemons); err != nil {
		return RespLocationWithPokemons{}, err
	}

	c.cache.Add(url, dat)

	return locationsWithPokemons, nil
}
