package providers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"providers_poc/cmd/api/domain"
)

var pokemonCache = make(map[string][]byte)

type MemCache struct {
	ID string
}

func NewMemCache() MemCache {
	return MemCache{
		ID: "memcache",
	}
}

type payloadMemCache struct {
	Name string `json:"name"`
}

func (m MemCache) GetId() string {
	return m.ID
}

func (m MemCache) RetrieveData(requestBody []byte) ([]byte, error) {
	p := payloadMemCache{}
	if err := json.Unmarshal(requestBody, &p); err != nil {
		panic(err)
	}

	pokemon, found := pokemonCache[p.Name]
	if !found {
		return nil, fmt.Errorf("pokemon not found in cache")
	}

	return pokemon, nil
}

func (m MemCache) AdaptData(b []byte) (domain.Pokemon, error) {
	var pokemon domain.Pokemon
	if err := json.Unmarshal(b, &pokemon); err != nil {
		return domain.Pokemon{}, err
	}

	pokemon.Provider = m.GetId()
	return pokemon, nil
}

func (m MemCache) WriteCache(pokemon domain.Pokemon) error {

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(pokemon)

	pokemonCache[pokemon.Name] = reqBodyBytes.Bytes()

	return nil
}
