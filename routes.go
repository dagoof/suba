package suba

import (
	"errors"
	"strings"
)

// Route functions based on strings as keys
type Keyed map[string]Handler

// Assign a handler function to a command
func (r Keyed) Assign(s string, f HF) { r[s] = HContainer{f} }

// Assign a sub-handler to a command
func (r Keyed) Sub(s string, h Handler) { r[s] = h }

// Returns the list of available commands for the given handler
func (r Keyed) Options() []string {
	keys := []string{}
	for key := range r {
		keys = append(keys, key)
	}
	return keys
}

// Delegate to handler function or sub-handler
func (r Keyed) Handle(args ...string) error {
	if len(args) > 0 {
		arg := args[0]
		args := args[1:]
		if h, ok := r[arg]; ok {
			return h.Handle(args...)
		}
	}
	return errors.New(strings.Join(r.Options(), "\n"))
}

// Route functions based on argument length.
// Useful for commands with optional arguments such as something like "git push"
type Counted map[int]HF

// Counted analogue to Keyed `Assign`
func (r Counted) Set(i int, f HF) { r[i] = f }

func (r Counted) Handle(args ...string) error {
	if f, ok := r[len(args)]; ok {
		return HContainer{f}.Handle(args...)
	}
	return INVALID_A
}

// Route both Keyed and Counted in one container. Makes Counted types easier to
// deal with by providing Keyed options as first choices, followed by Counted as
// fallbacks. 
// Instead of having a Counted for length 2 which needs to check the contents of
// the first against multiple keys, create a Keyed which captures these discrete
// cases, and a Counted for the fallbacks. See example
type Compound struct {
	Keyed
	Counted
}

func NewCompound() Compound {
	return Compound{Keyed{}, Counted{}}
}

func (c Compound) Handle(args ...string) error {
	if err := c.Keyed.Handle(args...); err == nil {
		return err
	}
	return c.Counted.Handle(args...)
}
