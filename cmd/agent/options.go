package main

import (
	"errors"
	"flag"
	"os"
	"strconv"
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
	ops := &options{}

	flag.StringVar(&ops.address, "a", _defaultAddress, "Address and port of server")
	flag.IntVar(&ops.reportInterval, "r", _defaultReportInterval, "Interval of sending reports, seconds")
	flag.IntVar(&ops.pollInterval, "p", _defaultPollInterval, "Interval of reading metrics, seconds")
	flag.Parse()

	if address, ok := os.LookupEnv("ADDRESS"); ok {
		ops.address = address
	}

	if reportInterval, ok := os.LookupEnv("REPORT_INTERVAL"); ok {
		i, err := strconv.ParseInt(reportInterval, 10, 32)
		if err != nil {
			return nil, errors.New("report interval value is not a number")
		}
		ops.reportInterval = int(i)
	}

	if pollInterval, ok := os.LookupEnv("POLL_INTERVAL"); ok {
		i, err := strconv.ParseInt(pollInterval, 10, 32)
		if err != nil {
			return nil, errors.New("poll interval value is not a number")
		}
		ops.pollInterval = int(i)
	}

	if ops.reportInterval <= 0 {
		return nil, errors.New("report interval too low or below zero")
	}

	if ops.pollInterval <= 0 {
		return nil, errors.New("poll interval too low or below zero")
	}

	if !strings.HasPrefix(ops.address, "http://") && !strings.HasPrefix(ops.address, "https://") {
		ops.address = "http://" + ops.address
	}

	return ops, nil
}
