package server

import (
	"github.com/Ceeler/metricssvc/internal/handler"
	"github.com/Ceeler/metricssvc/internal/repository"
	"github.com/Ceeler/metricssvc/internal/service"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := setupMetricRouter()
	router.Run("localhost:8080")
}

func setupMetricRouter() (router *gin.Engine) {
	router = gin.Default()
	router.RedirectTrailingSlash = false
	router.LoadHTMLFiles("internal/templates/index.tmpl")

	repo := repository.NewMemMetricRepository()
	svc := service.NewMetricService(repo)
	h := handler.NewMetricHandler(svc)
	router.POST("/update/:metric/:name/:value", h.UpdateHandler)
	router.GET("/value/:metric/:name", h.GetHandler)
	router.GET("/", h.GetAllHandler)
	return
}
