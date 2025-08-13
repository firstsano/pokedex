package main

import (
	"time"

	"github.com/firstsano/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &configuration{
		pokeClient: pokeClient,
	}

	startREPL(cfg)
}

/*
type Cache struct {
	mutex  *sync.Mutex
	values map[string]cacheEntry
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mutex:  &sync.Mutex{},
		values: make(map[string]cacheEntry),
	}
	ИЛИ
	cache := Cache{}

	return cache
}
*/
/*
var commandArguments []string равны ли
commandArguments = []string{}
*/
