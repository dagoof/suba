// Package suba - short for SUBcommand Argument parser - provides basic
// components useful in assembling a command line argument parser that
// uses a git-style subcommand based hierarchy of functions
package suba

import (
	"errors"
	"reflect"
)

var INVALID_A error = errors.New("Invalid argument count")
var INVALID_R error = errors.New("Invalid function result")

type Handler interface {
	Accept(...string) error
}

// Shorthand for `handler function`.
// Every function that eventually takes user input must have this
// function signature. Can be simplified through the use of 
// `Zero`, `One`, `Two`, and `Many` helper functions
type HF interface{}

// Simple container for HFs to allow for interface implementation
type HContainer struct{ F HF }

func (c HContainer) Accept(args ...string) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = INVALID_A
		}
	}()
	vargs := []reflect.Value{}
	for _, arg := range args {
		vargs = append(vargs, reflect.ValueOf(arg))
	}
	rs := reflect.ValueOf(c.F).Call(vargs)
	if len(rs) != 1 {
		return INVALID_R
	}
	r := rs[0]
	if r.Interface() == nil {
		return nil
	}
	return r.Interface().(error)
}

// Helper type that allows for function switching based on argument length count
// Useful for commands with optional arguments such as something like `git push`
type Many map[int]HF

func (m Many) Accept(args ...string) error {
	if f, ok := m[len(args)]; ok {
		return HContainer{ f }.Accept(args...)
	}
	return INVALID_A
}
