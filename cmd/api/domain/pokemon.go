package domain

import (
	"context"
	"errors"
)

const pokemonCtxKey ctxKey = "pokemonCtx"

type Pokemon struct {
	Provider string `json:"provider"`
	Name string `json:"name"`
	Type []string `json:"type"`
}


func AppendPokemonToCtx(ctx context.Context, pokemon Pokemon) context.Context {
	return context.WithValue(ctx, pokemonCtxKey, pokemon)
}

func GetPokemonFromCtx(ctx context.Context) (Pokemon, error) {
	if ctx == nil {
		return Pokemon{}, errors.New("empty context")
	}

	if pokemon, ok := ctx.Value(pokemonCtxKey).(Pokemon); ok {
		return pokemon, nil
	}

	return Pokemon{}, errors.New("pokemon not found in the context")
}