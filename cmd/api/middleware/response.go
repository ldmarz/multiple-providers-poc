package middleware

import (
	"encoding/json"
	"net/http"

	"providers_poc/cmd/api/domain"
)

func Response() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		pokemon, err := domain.GetPokemonFromCtx(r.Context())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pokemon)
	})
}