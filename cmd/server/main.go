package main

import "github.com/Ceeler/metricssvc/internal/server"

func main() {
	parseFlags()
	server.RunServer(*Address)
}
