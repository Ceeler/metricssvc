package main

import (
	"time"

	"github.com/Ceeler/metricssvc/internal/agent"
	"github.com/Ceeler/metricssvc/internal/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Initialize("info")
	log := logger.Get()
	ops, err := parseOptions()

	if err != nil {
		log.Fatal("Error during parse options", zap.Error(err))
	}

	a, err := agent.New(time.Duration(ops.pollInterval)*time.Second,
		time.Duration(ops.reportInterval)*time.Second,
		ops.address)

	if err != nil {
		log.Fatal("Error during starting agent", zap.Error(err))
	}

	a.Start()
}
