package suba

import "errors"

var INVALID_A error = errors.New("Invalid argument count")
type Handler interface {
	Accept(...string) error
}

type HF func(...string) error
type HContainer struct{ F HF }
func (c HContainer) Accept(args ...string) error { return c.F(args...) }

func Zero(f func() error) HF {
	g := func(args ...string) error {
		if len(args) == 0 {
			return f()
		}
		return INVALID_A
	}
	return g
}

func One(f func(string) error) HF {
	g := func(args ...string) error {
		if len(args) == 1 {
			return f(args[0])
		}
		return INVALID_A
	}
	return g
}

func Two(f func(string, string) error) HF {
	g := func(args ...string) error {
		if len(args) == 2 {
			return f(args[0], args[1])
		}
		return INVALID_A
	}
	return g
}

type Many map[int]HF
func (m Many) Accept(args ...string) error {
	if f, ok := m[len(args)]; ok {
		return f(args...)
	}
	return INVALID_A
}
