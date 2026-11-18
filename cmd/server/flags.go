package main

import "flag"

var (
	Address = flag.String("a", "localhost:8080", "Address and port of server")
)

func parseFlags() {
	flag.Parse()
}
