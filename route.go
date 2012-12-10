package suba

import (
	"strings"
	"errors"
)

type Route struct { Routes map[string]Handler }

func NewRoute() Route {
	return Route{map[string]Handler{}}
}

func (r Route) Add(s string, f HF) {
	r.Routes[s] = HContainer{ f }
}

func (r Route) Sub(s string, h Handler) {
	r.Routes[s] = h
}

func (r Route) Options() []string {
	keys := []string{}
	for key := range r.Routes {
		keys = append(keys, key)
	}
	return keys
}

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

