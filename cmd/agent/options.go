package main

import (
	"errors"
	"flag"
	"strings"
)

const (
	_defaultAddress        = "localhost:8080"
	_defaultReportInterval = 10
	_defaultPollInterval   = 2
)

type options struct {
	address        string
	reportInterval int
	pollInterval   int
}

func parseOptions() (*options, error) {
	ops := options{}
	flag.StringVar(&ops.address, "a", _defaultAddress, "Address and port of server")
	flag.IntVar(&ops.reportInterval, "r", _defaultReportInterval, "Interval of sending reports, seconds")
	flag.IntVar(&ops.pollInterval, "p", _defaultPollInterval, "Interval of reading metrics, seconds")

	if ops.reportInterval <= 0 {
		return nil, errors.New("report interval too low or below zero")
	}

	if ops.pollInterval <= 0 {
		return nil, errors.New("poll interval too low or below zero")
	}

	if !strings.HasPrefix(ops.address, "http://") && !strings.HasPrefix(ops.address, "https://") {
		ops.address = "http://" + ops.address
	}

	return &ops, nil
}
