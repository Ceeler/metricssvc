package handler

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/Ceeler/metricssvc/internal/model"
	"github.com/Ceeler/metricssvc/internal/service"
	"github.com/gin-gonic/gin"
)

type Metrics struct {
	service *service.Metric
}

func NewMetricHandler(svc *service.Metric) Metrics {
	return Metrics{
		service: svc,
	}
}

type modelDTO struct {
	Name  string
	Value string
}

func (h Metrics) GetAllHandler(c *gin.Context) {
	all := h.service.GetAllMetrics()
	dtos := make([]modelDTO, 0, len(all))
	for _, m := range all {
		switch t := models.MetricType(m.MType); t {
		case models.Counter:
			dtos = append(dtos, modelDTO{Name: string(m.ID), Value: fmt.Sprint(*m.Delta)})
		case models.Gauge:
			dtos = append(dtos, modelDTO{Name: string(m.ID), Value: fmt.Sprint(*m.Value)})
		}
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Metrics": dtos,
	})
}

func (h Metrics) GetHandler(c *gin.Context) {
	metricType := models.MetricType(c.Param("metric"))
	name := c.Param("name")
	metric := h.service.GetMetric(metricType, name)
	if metric != nil {
		var value string
		switch metricType {
		case models.Counter:
			value = fmt.Sprint(*metric.Delta)
		case models.Gauge:
			value = fmt.Sprint(*metric.Value)
		}
		c.String(http.StatusOK, "%s", value)
	} else {
		c.String(http.StatusNotFound, "Metric not found")
	}
}

func (h Metrics) UpdateHandler(c *gin.Context) {
	metricType := c.Param("metric")
	name := c.Param("name")
	valueStr := c.Param("value")

	if name == "" {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}

	if valueStr == "" {
		c.String(http.StatusBadRequest, "Value not presented")
		return
	}

	switch t := models.MetricType(metricType); t {
	case models.Counter:
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid counter value")
			return
		}
		h.service.UpdateCounter(name, value)
	case models.Gauge:
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid gauge value")
			return
		}
		h.service.UpdateGauge(name, value)
	default:
		c.String(http.StatusBadRequest, "Metric type not supported")
	}
}
