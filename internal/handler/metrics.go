package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ceeler/metricssvc/internal/logger"
	"github.com/Ceeler/metricssvc/internal/model"
	"github.com/Ceeler/metricssvc/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func (h Metrics) UpdateJsonHandler(c *gin.Context) {
	if c.ContentType() != gin.MIMEJSON {
		c.String(http.StatusUnsupportedMediaType, "Accepting json type")
	}
	var dto model.MetricsDto

	if err := json.NewDecoder(c.Request.Body).Decode(&dto); err != nil {
		logger.Get().Info("Error during decoding body", zap.Error(err))
		c.String(http.StatusBadRequest, "error decoding json body")
		return
	}

	switch t := model.MetricType(dto.MType); t {
	case model.Counter:
		if dto.Delta == nil {
			c.String(http.StatusBadRequest, "Delta null for Counter")
			return
		}
		h.service.UpdateCounter(dto.ID, *dto.Delta)
	case model.Gauge:
		if dto.Value == nil {
			c.String(http.StatusBadRequest, "Value null for Gauge")
			return
		}
		h.service.UpdateGauge(dto.ID, *dto.Value)
	default:
		c.String(http.StatusBadRequest, "Metric type not supported")
	}
}

func (h Metrics) GetAllJsonHandler(c *gin.Context) {
	all := h.service.GetAllMetrics()
	dtos := make([]model.MetricsDto, 0, len(all))
	for _, v := range all {
		dtos = append(dtos, model.MetricsDto{
			ID:    v.ID,
			MType: v.ID,
			Value: v.Value,
			Delta: v.Delta,
		})
	}
	if json, err := json.Marshal(dtos); err == nil {
		c.Header("Content-Type", gin.MIMEJSON)
		c.Writer.Write(json)
	} else {
		logger.Get().Error("Internal error encoding json", zap.Error(err))
		c.String(http.StatusInternalServerError, err.Error())
	}
}
