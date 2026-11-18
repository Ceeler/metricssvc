package service

import (
	"errors"

	models "github.com/Ceeler/metricssvc/internal/model"
	"github.com/Ceeler/metricssvc/internal/repository"
)

type MetricService struct {
	metricRepository repository.MetricRepository
}

func NewMetricService(repo repository.MetricRepository) *MetricService {
	return &MetricService{
		metricRepository: repo,
	}
}

func (ms *MetricService) GetMetric(metricType models.MetricType, name string) *models.Metrics {
	return ms.metricRepository.Get(metricType, name)
}

func (ms *MetricService) UpdateMetric(metricType models.MetricType, name string, value any) error {
	metric := ms.metricRepository.Get(metricType, name)

	if metric == nil {
		metric = &models.Metrics{
			ID:    name,
			MType: string(metricType),
		}
	}

	switch metricType {
	case models.Counter:
		var delta int64
		if metric.Delta == nil {
			delta = 0
		} else {
			delta = *metric.Delta
		}
		if v, ok := value.(int64); ok {
			delta += v
			metric.Delta = &delta
		} else {
			return errors.New("unsupported value of counter")
		}
	case models.Gauge:
		if v, ok := value.(float64); ok {
			metric.Value = &v
		} else {
			return errors.New("unsupported value of gauge")
		}
	default:
		return errors.New("unsupported metric type")
	}

	ms.metricRepository.Save(metric)
	return nil
}

func (ms MetricService) GetAllMetrics() []models.Metrics {
	return ms.metricRepository.GetAll()
}
