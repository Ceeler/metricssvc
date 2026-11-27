package repository

import (
	"github.com/Ceeler/metricssvc/internal/model"
	"github.com/Ceeler/metricssvc/internal/service"
)

type MemStorage struct {
	metricsMap map[model.MetricType]map[string]*model.Metrics
}

func NewMemStorage() *MemStorage {
	metricsMap := make(map[model.MetricType]map[string]*model.Metrics)
	return &MemStorage{
		metricsMap: metricsMap,
	}
}

var _ service.MetricStore = &MemStorage{}

// Получаем указатель, и кладем в мапу копию
func (m *MemStorage) Save(metric *model.Metrics) {
	metricType := model.MetricType(metric.MType)

	metrics, ok := m.metricsMap[metricType]
	if !ok {
		metrics = make(map[string]*model.Metrics)
		m.metricsMap[metricType] = metrics
	}

	metrics[metric.ID] = deepClone(metric)
}

// Возвращает ссылку на копию нужного элемента, или nil если элемент не найден
func (m *MemStorage) Get(metric model.MetricType, name string) *model.Metrics {
	metrics, ok := m.metricsMap[metric]
	if ok {
		return deepClone(metrics[name])
	}
	return nil
}

// Возвращает ссылки на копии элеметов из хранилища
func (m *MemStorage) GetAll() []*model.Metrics {
	all := make([]*model.Metrics, 0)
	for _, metrics := range m.metricsMap {
		for _, metric := range metrics {
			all = append(all, deepClone(metric))
		}
	}
	return all
}

func deepClone(original *model.Metrics) *model.Metrics {
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
