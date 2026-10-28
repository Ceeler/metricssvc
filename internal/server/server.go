package server

import (
	"net/http"

	"github.com/Ceeler/metricssvc/internal/handler"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/{metric}/{name}/{value}", handler.UpdateHandler)

	http.ListenAndServe("localhost:8080", mux)
}
