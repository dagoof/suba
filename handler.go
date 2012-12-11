// Package suba - short for SUBcommand Argument parser - provides basic
// components useful in assembling a command line argument parser that
// uses a git-style subcommand based hierarchy of functions
package suba

import "errors"

var INVALID_A error = errors.New("Invalid argument count")

type Handler interface {
	Accept(...string) error
}

// Shorthand for `handler function`.
// Every function that eventually takes user input must have this
// function signature. Can be simplified through the use of 
// `Zero`, `One`, `Two`, and `Many` helper functions
type HF func(...string) error

// Simple container for HFs to allow for interface implementation
type HContainer struct{ F HF }

func (c HContainer) Accept(args ...string) error { return c.F(args...) }

// Helper method that creates a HF from a zero-argument function
func Zero(f func() error) HF {
	g := func(args ...string) error {
		if len(args) == 0 {
			return f()
		}
		return INVALID_A
	}
	return g
}

// Helper method that creates a HF from a one-argument function
func One(f func(string) error) HF {
	g := func(args ...string) error {
		if len(args) == 1 {
			return f(args[0])
		}
		return INVALID_A
	}
	return g
}

// Helper method that creates a HF from a two-argument function
func Two(f func(string, string) error) HF {
	g := func(args ...string) error {
		if len(args) == 2 {
			return f(args[0], args[1])
		}
		return INVALID_A
	}
	return g
}

// Helper method that creates a HF from a three-argument function
func Three(f func(string, string, string) error) HF {
	g := func(args ...string) error {
		if len(args) == 3 {
			return f(args[0], args[1], args[2])
		}
		return INVALID_A
	}
	return g
}

// Helper type that allows for function switching based on argument length count
// Useful for commands with optional arguments such as something like `git push`
type Many map[int]HF

func (m Many) Accept(args ...string) error {
	if f, ok := m[len(args)]; ok {
		return f(args...)
	}
	return INVALID_A
}
