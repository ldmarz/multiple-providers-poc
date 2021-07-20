package middleware

import (
	"fmt"
	"net/http"

	"providers_poc/cmd/api/domain"
	"providers_poc/cmd/api/providers"
)

type Algo struct {

}

func SelectProvider(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		flow, ok := r.URL.Query()["flow"]
		if !ok || len(flow[0]) < 1 {
			http.Error(w, "Flow param is required", 400)
			return
		}

		provider, err := providerBasedOnFlow(flow[0])
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		ctx := domain.AppendProviderToCtx(r.Context(), provider)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}


func providerBasedOnFlow(flow string) ([]domain.Provider, error) {
	switch flow {
		case "pokemon":
			return []domain.Provider{providers.NewPokeApi()}, nil

		default:
			return nil, fmt.Errorf("flow not configured")
	}
}
