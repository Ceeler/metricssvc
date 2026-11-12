package main

import "github.com/Ceeler/metricssvc/internal/agent"

func main() {
	a := agent.NewAgent()
	a.Start()
}
