package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"providers_poc/cmd/api/domain"
)

const baseUrl = "https://pokeapi.co/api/v2/pokemon/"

type payloadPokeApi struct {
	Name string `json:"Name"`
}

type pokeApi struct {
	ID         string
	cacheStore WriteInCache
}

type responsePokeApi struct {
	Name string `json:"name"`
	Type []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func NewPokeApi(cacheStore WriteInCache) domain.Provider {
	return pokeApi{
		ID:         "pokeApi",
		cacheStore: cacheStore,
	}
}

func (pa pokeApi) GetId() string {
	return pa.ID
}

func (pa pokeApi) RetrieveData(requestBody []byte) ([]byte, error) {
	p := payloadPokeApi{}
	if err := json.Unmarshal(requestBody, &p); err != nil {
		return nil, err
	}

	resp, err := http.Get(baseUrl + p.Name)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pa pokeApi) AdaptData(b []byte) (domain.Pokemon, error) {
	var response responsePokeApi
	if err := json.Unmarshal(b, &response); err != nil {
		return domain.Pokemon{}, err
	}

	pokemon := domain.Pokemon{
		Provider: pa.GetId(),
		Name:     response.Name,
	}

	for _, v := range response.Type {
		pokemon.Type = append(pokemon.Type, v.Type.Name)
	}

	err := pa.cacheStore.WriteCache(pokemon)
	if err != nil {
		return domain.Pokemon{}, err
	}

	return pokemon, nil
}
