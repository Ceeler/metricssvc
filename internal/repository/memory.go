package repository

import (
	models "github.com/Ceeler/metricssvc/internal/model"
)

type MemStorage struct {
	metricsMap map[models.MetricType]map[string]*models.Metrics
}

func NewMemStorage() *MemStorage {
	metricsMap := make(map[models.MetricType]map[string]*models.Metrics)
	return &MemStorage{
		metricsMap: metricsMap,
	}
}

var _ Metric = &MemStorage{}

func (m *MemStorage) Save(metric models.Metrics) {
	metricType := models.MetricType(metric.MType)

	if existMetric, ok := m.Get(metricType, metric.ID); !ok {
		metrics, ok := m.metricsMap[metricType]
		if !ok {
			metrics = make(map[string]*models.Metrics)
			m.metricsMap[metricType] = metrics
		}
		metrics[metric.ID] = &metric
	} else {
		existMetric.Delta = metric.Delta
		existMetric.Value = metric.Value
		existMetric.Hash = metric.Hash
	}
}

func (m MemStorage) Get(metric models.MetricType, name string) (models.Metrics, bool) {
	metrics, ok := m.metricsMap[metric]
	if ok {
		m, ok := metrics[name]
		return *m, ok
	}
	return models.Metrics{}, false
}

func (m MemStorage) GetAll() []models.Metrics {
	all := make([]models.Metrics, 0)
	for _, metrics := range m.metricsMap {
		for _, metric := range metrics {
			all = append(all, *metric)
		}
	}
	return all
}
