package main

import (
	"flag"
	"os"
)

const _defaultAddress = "localhost:8080"

type options struct {
	address string
}

func parseOptions() *options {
	ops := &options{}

	flag.StringVar(&ops.address, "a", _defaultAddress, "Address and port of server")
	flag.Parse()

	if address, ok := os.LookupEnv("ADDRESS"); ok {
		ops.address = address
	}

	return ops
}
