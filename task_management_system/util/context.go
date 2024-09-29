package util

import (
	"context"
)

const RequestIDKey = "RequestIDKey"

type CustomContextValues struct {
	requestID string
}

// CustomContextKey struct
type CustomContextKey struct{}

func (custom CustomContextValues) GetRequestID(ctx *context.Context) string {
	return getContext(ctx).requestID
}

func getContext(ctx *context.Context) CustomContextValues {
	uContext, _ := (*ctx).Value(CustomContextKey{}).(CustomContextValues)
	return uContext
}
