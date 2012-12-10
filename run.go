package suba

import "os"

func Run(h Handler) error { return h.Accept(os.Args[1:]...) }
