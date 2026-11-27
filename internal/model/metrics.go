package model

type MetricType string

const (
	Counter MetricType = "counter"
	Gauge   MetricType = "gauge"
)

func (t MetricType) IsValid() bool {
	switch t {
	case Counter, Gauge:
		return true
	default:
		return false
	}
}

func (t MetricType) String() string {
	return string(t)
}

type Metrics struct {
	ID    string
	MType string
	Delta *int64
	Value *float64
	Hash  string
}

type MetricsDto struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}
