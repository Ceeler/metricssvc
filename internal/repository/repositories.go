package repository

import (
	models "github.com/Ceeler/metricssvc/internal/model"
)

type Metric interface {
	GetAll() []models.Metrics
	Get(metric models.MetricType, name string) (models.Metrics, bool)
	Save(models.Metrics)
}
