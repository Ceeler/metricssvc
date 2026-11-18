package main

import "github.com/Ceeler/metricssvc/internal/agent"

func main() {
	if err := parseFlags(); err != nil {
		panic(err)
	}

	a, err := agent.NewAgent(*PollInterval, *ReportInterval, *Address)
	if err != nil {
		panic(err)
	}

	a.Start()
}
