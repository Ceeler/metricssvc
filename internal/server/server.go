package server

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/Ceeler/metricssvc/internal/handler"
	"github.com/Ceeler/metricssvc/internal/logger"
	"github.com/Ceeler/metricssvc/internal/repository"
	"github.com/Ceeler/metricssvc/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:embed templates/*
var tmplFS embed.FS

func Run(address string) {
	router := gin.Default()
	router.Use(logger.Middleware)

	setupMetricRouter(router)
	if err := router.Run(address); err != nil {
		logger.Get().Fatal("Error during server start", zap.Error(err))
	}
}

func setupMetricRouter(router *gin.Engine) {
	router.RedirectTrailingSlash = false
	sub, _ := fs.Sub(tmplFS, "templates")
	router.LoadHTMLFS(http.FS(sub), "*.tmpl")

	repo := repository.NewMemStorage()
	svc := service.NewMetric(repo)
	h := handler.NewMetrics(svc)
	router.POST("/update/:metric/:name/:value", h.UpdateHandler)
	router.GET("/value/:metric/:name", h.GetHandler)
	router.GET("/", h.GetAllHandler)
	router.POST("/update", h.UpdateJsonHandler)
	router.GET("/value", h.GetAllJsonHandler)
}
