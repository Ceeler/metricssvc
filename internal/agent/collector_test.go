package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollector_CollectAppMetrics(t *testing.T) {
	c := &Collector{}
	am := c.CollectAppMetrics()

	assert.NotEmpty(t, am.Gauges, "Not empty gauge metrics")
	assert.NotEmpty(t, am.Counters, "Not empty counter metrics")

	assert.Len(t, am.Gauges, 28, "Should be 28 gauge metrics")
	assert.Len(t, am.Counters, 1, "Should be 1 counter metric")

	assert.Contains(t, am.Gauges, "Alloc")
	assert.Contains(t, am.Gauges, "BuckHashSys")
	assert.Contains(t, am.Gauges, "Frees")
	assert.Contains(t, am.Gauges, "GCCPUFraction")
	assert.Contains(t, am.Gauges, "GCSys")
	assert.Contains(t, am.Gauges, "HeapAlloc")
	assert.Contains(t, am.Gauges, "HeapIdle")
	assert.Contains(t, am.Gauges, "HeapInuse")
	assert.Contains(t, am.Gauges, "HeapObjects")
	assert.Contains(t, am.Gauges, "HeapReleased")
	assert.Contains(t, am.Gauges, "HeapSys")
	assert.Contains(t, am.Gauges, "LastGC")
	assert.Contains(t, am.Gauges, "Lookups")
	assert.Contains(t, am.Gauges, "MCacheInuse")
	assert.Contains(t, am.Gauges, "MCacheSys")
	assert.Contains(t, am.Gauges, "MSpanInuse")
	assert.Contains(t, am.Gauges, "MSpanSys")
	assert.Contains(t, am.Gauges, "Mallocs")
	assert.Contains(t, am.Gauges, "NextGC")
	assert.Contains(t, am.Gauges, "NumForcedGC")
	assert.Contains(t, am.Gauges, "NumGC")
	assert.Contains(t, am.Gauges, "OtherSys")
	assert.Contains(t, am.Gauges, "PauseTotalNs")
	assert.Contains(t, am.Gauges, "StackInuse")
	assert.Contains(t, am.Gauges, "StackSys")
	assert.Contains(t, am.Gauges, "Sys")
	assert.Contains(t, am.Gauges, "TotalAlloc")
	assert.Contains(t, am.Gauges, "RandomValue")
	assert.Contains(t, am.Counters, "PollCount")

	assert.Greater(t, am.Gauges["Sys"], 0.0)
	assert.Greater(t, am.Gauges["Alloc"], 0.0)

	assert.Equal(t, int64(1), am.Counters["PollCount"])

	c.CollectAppMetrics()
	am = c.CollectAppMetrics()

	assert.Equal(t, int64(3), am.Counters["PollCount"])
}
