package service

import (
	"github.com/Ceeler/metricssvc/internal/model"
)

type MetricStore interface {
	GetAll() []*model.Metrics
	Get(metric model.MetricType, name string) *model.Metrics
	Save(*model.Metrics)
}

type Metric struct {
	metricRepository MetricStore
}

func NewMetric(repo MetricStore) *Metric {
	return &Metric{
		metricRepository: repo,
	}
}

func (ms *Metric) GetMetric(metricType model.MetricType, name string) *model.Metrics {
	return ms.metricRepository.Get(metricType, name)
}

func (ms *Metric) UpdateCounter(name string, value int64) {
	metric := ms.metricRepository.Get(model.Counter, name)

	if metric == nil {
		metric = &model.Metrics{
			ID:    name,
			MType: string(model.Counter),
		}
	}

	if metric.Delta == nil {
		metric.Delta = new(int64)
	}

	*metric.Delta = *metric.Delta + value

	ms.metricRepository.Save(metric)
}

func (ms *Metric) UpdateGauge(name string, value float64) {
	metric := ms.metricRepository.Get(model.Gauge, name)

	if metric == nil {
		metric = &model.Metrics{
			ID:    name,
			MType: string(model.Gauge),
		}
	}

	metric.Value = &value

	ms.metricRepository.Save(metric)
}

func (ms *Metric) GetAllMetrics() []*model.Metrics {
	return ms.metricRepository.GetAll()
}
