package handler

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/Ceeler/metricssvc/internal/model"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	metricType, name, valueStr := r.PathValue("metric"), r.PathValue("name"), r.PathValue("value")

	if name == "" {
		http.NotFound(w, r)
		return
	}

	if valueStr == "" {
		http.Error(w, "Value type not supported", http.StatusBadRequest)
		return
	}

	switch metricType {
	case models.Counter:
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid counter value", http.StatusBadRequest)
			return
		}
		fmt.Printf("Metric: %s Name: %s Value %v", metricType, name, value)
	case models.Gauge:
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			http.Error(w, "Invalid counter value", http.StatusBadRequest)
			return
		}
		fmt.Printf("Metric: %s Name: %s Value %v", metricType, name, value)
	default:
		http.Error(w, "Metric type not supported", http.StatusBadRequest)
	}
}
