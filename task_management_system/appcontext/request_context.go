package appcontext

import (
	"context"
)

// RequestContextKey is a request context key
type RequestContextKey = struct{}

type RequestContext struct {
	tenant    string
	requestID string
}

func (r *RequestContext) Tenant() string {
	if r == nil {
		return ""
	}
	return r.tenant
}

func (r *RequestContext) RequestID() string {
	if r == nil {
		return ""
	}
	return r.requestID
}

func NewRequestContext(tenant, requestID string) *RequestContext {
	return &RequestContext{
		tenant:    tenant,
		requestID: requestID,
	}
}

func AddRequestContext(ctx context.Context, requestContext *RequestContext) context.Context {
	return context.WithValue(ctx, RequestContextKey{}, requestContext)
}

func GetRequestContext(ctx context.Context) *RequestContext {
	requestContext, _ := ctx.Value(RequestContextKey{}).(*RequestContext)
	return requestContext
}

func AddOfflineContext(ctx context.Context, tenantID, requestID string) context.Context {
	reqContext := NewRequestContext(tenantID, requestID)
	return context.WithValue(ctx, RequestContextKey{}, reqContext)
}

//Copy returns a context.Background() with RequestContext value
func Copy(source context.Context) context.Context {
	reqContext := GetRequestContext(source)
	newContext := AddRequestContext(context.Background(), reqContext)
	return newContext
}

type TransactionId struct{}
