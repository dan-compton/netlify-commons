package tracing

import (
	"context"
	"net/http"
)

type contextKey string

const tracerKey = contextKey("nf-tracer-key")

func WrapWithTracer(r *http.Request, rt *RequestTracer) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), tracerKey, rt))
}

func GetTracer(r *http.Request) *RequestTracer {
	val := r.Context().Value(tracerKey)
	if val == nil {
		return nil
	}
	entry, ok := val.(*RequestTracer)
	if ok {
		return entry
	}
	return nil
}