package main

import (
	"context"

	"github.com/cuonglm/gogi"
)

type contextKey string

var clientKey = contextKey("client")

// WithClient adds a gogi HTTP client to the context.
func WithClient(ctx context.Context, client *gogi.Client) context.Context {
	return context.WithValue(ctx, clientKey, client)
}

// ClientFromContext retrieves the gogi client from the context.
func ClientFromContext(ctx context.Context) (*gogi.Client, bool) {
	client, ok := ctx.Value(clientKey).(*gogi.Client)
	return client, ok
}
