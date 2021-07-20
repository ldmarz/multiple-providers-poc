package domain

import (
	"context"
	"errors"
	"io"
)

const providerCtxKey ctxKey = "providerCtx"

type Provider interface {
	GetId() string
	RetrieveData(requestBody io.ReadCloser) ([]byte, error)
	AdaptData(data []byte) (Pokemon, error)
}

func AppendProviderToCtx(ctx context.Context, provider []Provider) context.Context {
	return context.WithValue(ctx, providerCtxKey, provider)
}

func GetProviderFromCtx(ctx context.Context) ([]Provider, error) {
	if ctx == nil {
		return nil, errors.New("empty context")
	}

	if provider, ok := ctx.Value(providerCtxKey).([]Provider); ok {
		return provider, nil
	}

	return nil, errors.New("provider not found in the context")
}