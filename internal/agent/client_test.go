package agent

import (
	"net/http"
	"testing"

	"io"
	"net/http/httptest"

	models "github.com/Ceeler/metricssvc/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpSender_WrongUrl(t *testing.T) {
	_, err := NewDefaultSender("://myServer")
	require.Error(t, err)
}

func TestHttpSender_SendWithNoServer(t *testing.T) {
	sender, err := NewDefaultSender("http://localhost:8080/")
	require.NoError(t, err)

	err = sender.SendMetric(models.Gauge, "Alloc", "123.5")

	require.Error(t, err)
}

func TestHttpSender_SendMetric(t *testing.T) {
	received := make(chan struct{}, 1)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/update/gauge/Alloc/123.5", r.URL.Path)
			assert.Equal(t, "text/plain", r.Header.Get("Content-Type"))

			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.Empty(t, body)

			received <- struct{}{}
			w.WriteHeader(http.StatusOK)
		},
	))
	defer server.Close()

	sender, err := NewDefaultSender(server.URL)
	require.NoError(t, err)

	err = sender.SendMetric(models.Gauge, "Alloc", "123.5")

	require.NoError(t, err)

	require.Len(t, received, 1, "Server should receive metric")
}

func TestHttpSender_SendErrorMetric(t *testing.T) {
	received := make(chan struct{}, 1)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/update/gauge/Alloc/123.5", r.URL.Path)
			assert.Equal(t, "text/plain", r.Header.Get("Content-Type"))

			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.Empty(t, body)

			received <- struct{}{}
			w.WriteHeader(http.StatusBadRequest)
		},
	))
	defer server.Close()

	sender, err := NewDefaultSender(server.URL)
	require.NoError(t, err)

	err = sender.SendMetric(models.Gauge, "Alloc", "123.5")

	require.Error(t, err)

	require.Len(t, received, 1, "Server should receive metric")
}
