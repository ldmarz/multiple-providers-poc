package providers

import "providers_poc/cmd/api/domain"

type WriteInCache interface {
	WriteCache(pokemon domain.Pokemon) error
}
