package repository

import (
	models "github.com/Ceeler/metricssvc/internal/model"
)

type MetricRepository interface {
	GetAll() []models.Metrics
	Get(metric models.MetricType, name string) *models.Metrics
	Save(*models.Metrics)
}
