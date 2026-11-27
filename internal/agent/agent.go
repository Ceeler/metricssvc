package agent

import (
	"sync"
	"time"

	"github.com/Ceeler/metricssvc/internal/logger"
	"github.com/Ceeler/metricssvc/internal/model"
	"go.uber.org/zap"
)

type Agent struct {
	PollInterval   time.Duration
	ReportInterval time.Duration
	Collector      MetricCollector
	Sender         MetricSender
}

func New(pollInterval, reportInterval time.Duration, address string) (*Agent, error) {
	sender, err := NewDefaultSender(address)
	if err != nil {
		return nil, err
	}

	return &Agent{
		PollInterval:   pollInterval,
		ReportInterval: reportInterval,
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
			err := a.Sender.SendMetricJSON(model.Gauge, metric, nil, &value)
			auditError(err)
		}

		for metric, value := range counters {
			err := a.Sender.SendMetricJSON(model.Counter, metric, &value, nil)
			auditError(err)
		}

	}
}

func auditError(err error) {
	if err != nil {
		logger.Get().Warn("Error during sending metric to server", zap.Error(err))
	}
}
