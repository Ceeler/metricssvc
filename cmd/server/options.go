package main

import "flag"

type options struct {
	address string
}

func parseOptions() *options {
	ops := options{}
	flag.StringVar(&ops.address, "a", "localhost:8080", "Address and port of server")
	flag.Parse()

	return &ops
}
