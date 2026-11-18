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

// Получаем указатель, и кладем в мапу копию
func (m *MemStorage) Save(metric *models.Metrics) {
	metricType := models.MetricType(metric.MType)

	metrics, ok := m.metricsMap[metricType]
	if !ok {
		metrics = make(map[string]*models.Metrics)
		m.metricsMap[metricType] = metrics
	}

	metrics[metric.ID] = deepClone(metric)
}

// Возвращает ссылку на копию нужного элемента, или nil если элемент не найден
func (m *MemStorage) Get(metric models.MetricType, name string) *models.Metrics {
	metrics, ok := m.metricsMap[metric]
	if ok {
		return deepClone(metrics[name])
	}
	return nil
}

// Возвращает ссылки на копии элеметов из хранилища
func (m *MemStorage) GetAll() []*models.Metrics {
	all := make([]*models.Metrics, 0)
	for _, metrics := range m.metricsMap {
		for _, metric := range metrics {
			all = append(all, deepClone(metric))
		}
	}
	return all
}

func deepClone(original *models.Metrics) *models.Metrics {
	if original == nil {
		return nil
	}
	copy := *original
	if original.Delta != nil {
		d := *original.Delta
		copy.Delta = &d
	}
	if original.Value != nil {
		v := *original.Value
		copy.Value = &v
	}
	return &copy
}
