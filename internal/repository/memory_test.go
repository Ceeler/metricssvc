package repository

import (
	"testing"

	models "github.com/Ceeler/metricssvc/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemStorage_Save(t *testing.T) {
	storage := NewMemStorage()
	startValue := 10.0
	testMetric := models.Metrics{
		ID:    "test",
		MType: string(models.Gauge),
		Value: new(startValue),
	}
	storage.Save(&testMetric)

	all := storage.GetAll()
	require.Equal(t, 1, len(all), "One value saved")

	savedMetric := storage.metricsMap[models.MetricType(testMetric.MType)][testMetric.ID]

	*testMetric.Value = 50
	assert.Equal(t, startValue, *savedMetric.Value, "Test that saved metric not connected with test metric")

	allMetric := all[0]
	*allMetric.Value = 60
	assert.Equal(t, startValue, *savedMetric.Value, "Test that saved metric not connected with returned all metrics")

	getMetric := storage.Get(models.Gauge, testMetric.ID)
	require.NotNil(t, getMetric, "Found saved metric")

	*getMetric.Value = 70
	assert.Equal(t, startValue, *savedMetric.Value, "Test that saved metric not connected with returned all metrics")

	storage.Save(&testMetric)
	getNewMetric := storage.Get(models.Gauge, testMetric.ID)
	assert.Equal(t, *testMetric.Value, *getNewMetric.Value, "Test that saved update saved metric")

}
