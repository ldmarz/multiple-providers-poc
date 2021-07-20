package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"providers_poc/cmd/api/domain"
)

func RetrieveData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providers, err := domain.GetProviderFromCtx(r.Context())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var pokemon domain.Pokemon
		body, err := ioutil.ReadAll(r.Body)
		for _, provider := range providers {
			rawData, err := provider.RetrieveData(body)
			if err != nil {
				fmt.Println(fmt.Errorf("some error occurs using the provider %s, error: %w", provider.GetId(), err))
				continue
			}

			pokemon, err = provider.AdaptData(rawData)
			if err != nil {
				fmt.Println(fmt.Errorf("some error occurs parsing the data from provider %s, error: %w", provider.GetId(), err))
				continue
			}

			// If not errors occurs the data should be fine
			break
		}

		ctx := domain.AppendPokemonToCtx(r.Context(), pokemon)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
