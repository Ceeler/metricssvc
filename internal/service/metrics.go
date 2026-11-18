package service

import (
	models "github.com/Ceeler/metricssvc/internal/model"
	"github.com/Ceeler/metricssvc/internal/repository"
)

type Metric struct {
	metricRepository repository.Metric
}

func NewMetricService(repo repository.Metric) *Metric {
	return &Metric{
		metricRepository: repo,
	}
}

func (ms *Metric) GetMetric(metricType models.MetricType, name string) *models.Metrics {
	return ms.metricRepository.Get(metricType, name)
}

func (ms *Metric) UpdateCounter(name string, value int64) {
	metric := ms.metricRepository.Get(models.Counter, name)

	if metric == nil {
		metric = &models.Metrics{
			ID:    name,
			MType: string(models.Counter),
		}
	}

	if metric.Delta == nil {
		metric.Delta = new(int64)
	}

	*metric.Delta = *metric.Delta + value

	ms.metricRepository.Save(metric)
}

func (ms *Metric) UpdateGauge(name string, value float64) {
	metric := ms.metricRepository.Get(models.Gauge, name)

	if metric == nil {
		metric = &models.Metrics{
			ID:    name,
			MType: string(models.Gauge),
		}
	}

	metric.Value = &value

	ms.metricRepository.Save(metric)
}

func (ms *Metric) GetAllMetrics() []*models.Metrics {
	return ms.metricRepository.GetAll()
}
