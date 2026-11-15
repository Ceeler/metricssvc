package repository

import (
	models "github.com/Ceeler/metricssvc/internal/model"
)

type MemMetricRepository struct {
	metricsMap map[models.MetricType]map[string]*models.Metrics
}

func NewMemMetricRepository() *MemMetricRepository {
	metricsMap := make(map[models.MetricType]map[string]*models.Metrics)
	return &MemMetricRepository{
		metricsMap: metricsMap,
	}
}

var _ MetricRepository = &MemMetricRepository{}

func (m *MemMetricRepository) Save(metric *models.Metrics) {
	metricType := models.MetricType(metric.MType)
	if existMetric := m.Get(metricType, metric.ID); existMetric == nil {
		metrics, ok := m.metricsMap[metricType]
		if !ok {
			metrics = make(map[string]*models.Metrics)
			m.metricsMap[metricType] = metrics
		}
		metrics[metric.ID] = metric
	} else {
		existMetric.Delta = metric.Delta
		existMetric.Value = metric.Value
		existMetric.Hash = metric.Hash
	}
}

func (m MemMetricRepository) Get(metric models.MetricType, name string) *models.Metrics {
	metrics, ok := m.metricsMap[metric]
	if ok {
		return metrics[name]
	}
	return nil
}

func (m MemMetricRepository) GetAll() []models.Metrics {
	all := make([]models.Metrics, 0)
	for _, metrics := range m.metricsMap {
		for _, metric := range metrics {
			all = append(all, *metric)
		}
	}
	return all
}
