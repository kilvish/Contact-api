package middleware

import (
	"net/http"
)

// ContextKey is used to generate context keys for middlewares
type ContextKey string

// Middleware represents a middleware
type Middleware func(http.Handler) http.Handler

// Chain represente the chain of middlewares
type Chain struct {
	middlewares []Middleware
}

// New creates a new chain of middlewares
func New(middlewares ...Middleware) *Chain {
	return &Chain{append(([]Middleware)(nil), middlewares...)}
}

// Then fucntion adds the handler at the end of the middlewares
func (c *Chain) Then(handler http.Handler) http.Handler {
	for i := range c.middlewares {
		handler = c.middlewares[len(c.middlewares)-i-1](handler)
	}
	return handler
}

// Add adds the middleware
func (c *Chain) Add(middlewares ...Middleware) *Chain {
	newMiddlewares := make([]Middleware, 0)
	newMiddlewares = append(newMiddlewares, c.middlewares...)
	newMiddlewares = append(newMiddlewares, middlewares...)
	return &Chain{newMiddlewares}
}

