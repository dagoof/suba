package suba

import (
	"errors"
	"strings"
)

type Route struct{ Routes map[string]Handler }

// Create an empty route
func NewRoute() Route {
	return Route{map[string]Handler{}}
}

// Handle the given command with a handler function
func (r Route) Add(s string, f HF) {
	r.Routes[s] = HContainer{f}
}

// Handle the given command with a sub-handler
func (r Route) Sub(s string, h Handler) {
	r.Routes[s] = h
}

// Returns the list of available commands for the given handler
func (r Route) Options() []string {
	keys := []string{}
	for key := range r.Routes {
		keys = append(keys, key)
	}
	return keys
}

// Method implementing `handler` interface by delegating to handler function or
// sub-handler
func (r Route) Accept(args ...string) error {
	if len(args) > 0 {
		arg := args[0]
		args := args[1:]
		if h, ok := r.Routes[arg]; ok {
			return h.Accept(args...)
		}
	}
	return errors.New(strings.Join(r.Options(), "\n"))
}
