package agent

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Ceeler/metricssvc/internal/model"
)

type MetricSender interface {
	SendMetric(metric model.MetricType, name, value string) error
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

func (hs HTTPSender) SendMetric(metric model.MetricType, name, value string) error {
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
