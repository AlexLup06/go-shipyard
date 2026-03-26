package httpctx

import "context"

type htmxKeyType struct{}

var (
	htmxKey = htmxKeyType{}
)

func WithHTMX(ctx context.Context) context.Context {
	return context.WithValue(ctx, htmxKey, true)
}

func IsHTMX(ctx context.Context) bool {
	v, ok := ctx.Value(htmxKey).(bool)
	return ok && v
}
