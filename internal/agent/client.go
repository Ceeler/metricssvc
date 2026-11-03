package agent

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	models "github.com/Ceeler/metricssvc/internal/model"
)

type MetricSender interface {
	SendMetric(metric models.MetricType, name, value string) error
}

type HttpSender struct {
	Client *http.Client
	URL    *url.URL
}

var _ MetricSender = HttpSender{}

func NewDefaultSender(host string) (*HttpSender, error) {
	url, err := url.Parse(host)

	if err != nil {
		return nil, err

	}

	client := &http.Client{}
	client.Timeout = time.Duration(5) * time.Second

	return &HttpSender{
		Client: client,
		URL:    url,
	}, nil
}

func (hs HttpSender) SendMetric(metric models.MetricType, name, value string) error {
	metricUrl := hs.URL.JoinPath("update", string(metric), name, value).String()
	resp, err := hs.Client.Post(metricUrl, "text/plain", nil)

	if err == nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		if resp.StatusCode >= 300 {
			return fmt.Errorf("Server return not success status: %d", resp.StatusCode)
		}
	}

	return err
}
