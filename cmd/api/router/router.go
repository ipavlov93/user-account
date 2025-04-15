package router

import (
	"net/http"
)

// Middleware type as before
type Middleware func(http.Handler) http.Handler

// Router struct to hold our routes and middleware
type Router struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

// NewRouter creates and returns a new Router with an initialized ServeMux and middleware slice
func NewRouter() *Router {
	return &Router{
		mux:         http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

// Use adds middleware to the chain
func (a *Router) Use(mw Middleware) {
	a.middlewares = append(a.middlewares, mw)
}

// Handle registers a handler for a specific route, applying all middleware
func (a *Router) Handle(pattern string, handler http.Handler) {
	finalHandler := handler
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		finalHandler = a.middlewares[i](finalHandler)
	}
	a.mux.Handle(pattern, finalHandler)
}

// ListenAndServe starts the application server
func (a *Router) ListenAndServe(address string) error {
	return http.ListenAndServe(address, a.mux)
}
