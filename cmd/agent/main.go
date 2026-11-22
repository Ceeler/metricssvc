package main

import (
	"time"

	"github.com/Ceeler/metricssvc/internal/agent"
)

func main() {
	ops, err := parseOptions()

	if err != nil {
		panic(err)
	}

	a, err := agent.New(time.Duration(ops.pollInterval)*time.Second,
		time.Duration(ops.reportInterval)*time.Second,
		ops.address)

	if err != nil {
		panic(err)
	}

	a.Start()
}
