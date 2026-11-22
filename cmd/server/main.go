package main

import (
	"github.com/Ceeler/metricssvc/internal/logger"
	"github.com/Ceeler/metricssvc/internal/server"
)

func main() {
	logger.Initialize("info")
	ops := parseOptions()
	server.Run(ops.address)
}
