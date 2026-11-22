package main

import "github.com/Ceeler/metricssvc/internal/server"

func main() {
	ops := parseOptions()
	server.Run(ops.address)
}
