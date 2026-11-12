package agent

import (
	"math/rand"
	"reflect"
	"runtime"
)

type AppMetrics struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

type MetricCollector interface {
	CollectAppMetrics() AppMetrics
}

type Collector struct {
	pollCounter int64
}

var _ MetricCollector = &Collector{}

func (c *Collector) CollectAppMetrics() (am AppMetrics) {
	c.pollCounter++
	am = AppMetrics{
		Gauges:   make(map[string]float64),
		Counters: make(map[string]int64),
	}

	c.collectRuntime(&am)
	c.collectCustom(&am)

	return
}

func (c *Collector) collectRuntime(am *AppMetrics) {
	gauges := am.Gauges

	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	for field, value := range reflect.ValueOf(stats).Fields() {
		if value.CanUint() {
			gauges[field.Name] = float64(value.Uint())
		}
		if value.CanFloat() {
			gauges[field.Name] = value.Float()
		}
	}
}

func (c *Collector) collectCustom(am *AppMetrics) {
	am.Gauges["RandomValue"] = rand.Float64()
	am.Counters["PollCount"] = c.pollCounter
}
