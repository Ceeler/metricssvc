package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ceeler/metricssvc/internal/model"
	"github.com/Ceeler/metricssvc/internal/service"
	"github.com/gin-gonic/gin"
)

type Metrics struct {
	service *service.Metric
}

func NewMetrics(svc *service.Metric) Metrics {
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
		switch t := model.MetricType(m.MType); t {
		case model.Counter:
			dtos = append(dtos, modelDTO{Name: string(m.ID), Value: fmt.Sprint(*m.Delta)})
		case model.Gauge:
			dtos = append(dtos, modelDTO{Name: string(m.ID), Value: fmt.Sprint(*m.Value)})
		}
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Metrics": dtos,
	})
}

func (h Metrics) GetHandler(c *gin.Context) {
	metricType := model.MetricType(c.Param("metric"))
	name := c.Param("name")
	metric := h.service.GetMetric(metricType, name)
	if metric == nil {
		c.String(http.StatusNotFound, "Metric not found")
		return
	}

	var value string
	switch metricType {
	case model.Counter:
		value = fmt.Sprint(*metric.Delta)
	case model.Gauge:
		value = fmt.Sprint(*metric.Value)
	}
	c.String(http.StatusOK, "%s", value)
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

	switch t := model.MetricType(metricType); t {
	case model.Counter:
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid counter value")
			return
		}
		h.service.UpdateCounter(name, value)
	case model.Gauge:
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
