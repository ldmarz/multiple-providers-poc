package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"providers_poc/cmd/api/domain"
)

const baseUrlPokeXYZ = "https://app.pokemon-api.xyz/pokemon/"

type pokeXYZ struct {
	ID         string
	cacheStore WriteInCache
}

type payloadPokeXYZ struct {
	Name string `json:"name"`
}

type responseXYZ struct {
	Name struct {
		English string `json:"english"`
	} `json:"name"`
	Type []string `json:"type"`
}

func NewPokeXYZ(cacheStore WriteInCache) domain.Provider {
	return pokeXYZ{
		ID:         "pokeXYZ",
		cacheStore: cacheStore,
	}
}

func (p pokeXYZ) GetId() string {
	return p.ID
}

func (pxyz pokeXYZ) RetrieveData(requestBody []byte) ([]byte, error) {
	p := payloadPokeXYZ{}
	if err := json.Unmarshal(requestBody, &p); err != nil {
		return nil, err
	}

	resp, err := http.Get(baseUrlPokeXYZ + p.Name)
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

func (pxyz pokeXYZ) AdaptData(b []byte) (domain.Pokemon, error) {
	var response responseXYZ
	if err := json.Unmarshal(b, &response); err != nil {
		return domain.Pokemon{}, err
	}

	pokemon := domain.Pokemon{
		Provider: pxyz.GetId(),
		Name:     response.Name.English,
	}

	for _, v := range response.Type {
		pokemon.Type = append(pokemon.Type, v)
	}

	err := pxyz.cacheStore.WriteCache(pokemon)
	if err != nil {
		return domain.Pokemon{}, err
	}

	return pokemon, nil
}
