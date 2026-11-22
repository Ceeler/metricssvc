package server

import (
	"github.com/Ceeler/metricssvc/internal/handler"
	"github.com/Ceeler/metricssvc/internal/repository"
	"github.com/Ceeler/metricssvc/internal/service"
	"github.com/gin-gonic/gin"
)

func Run(address string) {
	router := setupMetricRouter()
	if err := router.Run(address); err != nil {
		panic(err)
	}
}

func setupMetricRouter() (router *gin.Engine) {
	router = gin.Default()
	router.RedirectTrailingSlash = false
	router.LoadHTMLFiles("internal/templates/index.tmpl")

	repo := repository.NewMemStorage()
	svc := service.NewMetric(repo)
	h := handler.NewMetrics(svc)
	router.POST("/update/:metric/:name/:value", h.UpdateHandler)
	router.GET("/value/:metric/:name", h.GetHandler)
	router.GET("/", h.GetAllHandler)
	return
}
