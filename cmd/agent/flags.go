package main

import (
	"errors"
	"flag"
	"strings"
)

var (
	Address        = flag.String("a", "localhost:8080", "Address and port of server")
	ReportInterval = flag.Int("r", 10, "Interval of sending reports, seconds")
	PollInterval   = flag.Int("p", 2, "Interval of reading metrics, seconds")
)

func parseFlags() error {
	flag.Parse()
	if *ReportInterval <= 0 {
		return errors.New("report interval too low or below zero")
	}

	if *PollInterval <= 0 {
		return errors.New("poll interval too low or below zero")
	}

	if !strings.HasPrefix(*Address, "http://") && !strings.HasPrefix(*Address, "https://") {
		*Address = "http://" + *Address
	}

	return nil
}
