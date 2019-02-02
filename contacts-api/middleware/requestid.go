package middleware

import (
	"context"
	"net/http"

	"github.com/anil/loomtest/src/uuid"
)

const (
	// RequestIDContextKey - context key name for reqeust ID
	RequestIDContextKey = "requestID"
)

// GetRequestIDMiddleware gets the middleware to generate the request ID
func GetRequestIDMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.GetUUID()
			r = r.WithContext(context.WithValue(r.Context(), ContextKey(RequestIDContextKey), requestID))
			h.ServeHTTP(w, r)
		})
	}

}

