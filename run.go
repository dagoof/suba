package suba

import "os"

// Run a handler on default command line args
func Run(h Handler) error { return h.Accept(os.Args[1:]...) }
