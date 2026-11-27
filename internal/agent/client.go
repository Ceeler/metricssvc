package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Ceeler/metricssvc/internal/logger"
	"github.com/Ceeler/metricssvc/internal/model"
	"go.uber.org/zap"
)

type MetricSender interface {
	SendMetric(metric model.MetricType, name, value string) error
	SendMetricJSON(metric model.MetricType, name string, delta *int64, value *float64) error
}

type HTTPSender struct {
	Client *http.Client
	URL    *url.URL
}

var _ MetricSender = HTTPSender{}

func NewDefaultSender(host string) (*HTTPSender, error) {
	url, err := url.Parse(host)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	client.Timeout = time.Duration(5) * time.Second

	return &HTTPSender{
		Client: client,
		URL:    url,
	}, nil
}

// Deprecated use SendMetricJson instead
func (hs HTTPSender) SendMetric(metric model.MetricType, name, value string) error {
	logger.Get().Info("Send metric", zap.String("type", string(metric)),
		zap.String("name", name),
		zap.String("value", value))

	metricURL := hs.URL.JoinPath("update", string(metric), name, value).String()
	resp, err := hs.Client.Post(metricURL, "text/plain", nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("server return not success status: %d", resp.StatusCode)
	}

	return nil
}

func (hs HTTPSender) SendMetricJSON(metric model.MetricType, name string, delta *int64, value *float64) error {
	metricURL := hs.URL.JoinPath("update").String()
	dto := model.MetricsDto{
		ID:    name,
		MType: string(metric),
		Value: value,
		Delta: delta,
	}
	body, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	resp, err := hs.Client.Post(metricURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("server return not success status: %d", resp.StatusCode)
	}

	return nil
}
