package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerUpdateHandler(t *testing.T) {
	type want struct {
		statusCode int
		message    string
	}
	type args struct {
		url, method string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test correct gauge",
			args: args{
				method: http.MethodPost,
				url:    "/update/gauge/Test/125.2",
			},
			want: want{
				statusCode: 200,
				message:    "",
			},
		},
		{
			name: "Test correct counter",
			args: args{
				method: http.MethodPost,
				url:    "/update/counter/Test/125",
			},
			want: want{
				statusCode: 200,
				message:    "",
			},
		},
		{
			name: "Test not correct method",
			args: args{
				method: http.MethodPut,
				url:    "/update/gauge/Test/125.2",
			},
			want: want{
				statusCode: 404,
				message:    "404 page not found",
			},
		},
		{
			name: "Test not correct metric",
			args: args{
				method: http.MethodPost,
				url:    "/update/timer/Test/125.2",
			},
			want: want{
				statusCode: 400,
				message:    "Metric type not supported",
			},
		},
		{
			name: "Test not correct gauge value",
			args: args{
				method: http.MethodPost,
				url:    "/update/gauge/Test/PT15S",
			},
			want: want{
				statusCode: 400,
				message:    "Invalid gauge value",
			},
		},
		{
			name: "Test not correct counter value",
			args: args{
				method: http.MethodPost,
				url:    "/update/counter/Test/125.2",
			},
			want: want{
				statusCode: 400,
				message:    "Invalid counter value",
			},
		},
		{
			name: "Test not correct path",
			args: args{
				method: http.MethodPost,
				url:    "/update/counter/",
			},
			want: want{
				statusCode: 404,
				message:    "404 page not found",
			},
		},
	}

	router := gin.Default()
	setupMetricRouter(router)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequestWithContext(t.Context(), tt.args.method, tt.args.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)

			result := w.Result()

			userResult, err := io.ReadAll(result.Body)
			body := string(userResult)

			require.Nil(t, err, "Incorrect body")
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.message, body)
		})
	}
}
