package agent

import (
	"fmt"
	"sync"
	"time"

	models "github.com/Ceeler/metricssvc/internal/model"
)

type Agent struct {
	PollInterval   time.Duration
	ReportInterval time.Duration
	Collector      MetricCollector
	Sender         MetricSender
}

func NewAgent(pollInterval, reportInterval int, address string) (*Agent, error) {
	sender, err := NewDefaultSender(address)
	if err != nil {
		return nil, err
	}

	return &Agent{
		PollInterval:   time.Duration(pollInterval) * time.Second,
		ReportInterval: time.Duration(reportInterval) * time.Second,
		Collector:      &Collector{},
		Sender:         sender,
	}, nil
}

func (a *Agent) Start() {

	appMetrics := AppMetrics{}
	mux := sync.Mutex{}

	go func() {
		for {
			time.Sleep(a.PollInterval)
			mux.Lock()
			appMetrics = a.Collector.CollectAppMetrics()
			mux.Unlock()
		}
	}()

	for {
		time.Sleep(a.ReportInterval)
		mux.Lock()
		gauges := appMetrics.Gauges
		counters := appMetrics.Counters
		mux.Unlock()

		for metric, value := range gauges {
			err := a.Sender.SendMetric(models.Gauge, metric, fmt.Sprint(value))
			auditError(err)
		}

		for metric, value := range counters {
			err := a.Sender.SendMetric(models.Counter, metric, fmt.Sprint(value))
			auditError(err)
		}

	}
}

func auditError(err error) {
	if err != nil {
		fmt.Printf("Error during sending metric to server %v", err)
	}
}
