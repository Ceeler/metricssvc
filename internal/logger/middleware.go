package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Middleware(c *gin.Context) {
	t := time.Now()

	c.Next()

	latency := time.Since(t)
	log.Info("Receive request",
		zap.String("Method", c.Request.Method),
		zap.String("URI", c.Request.URL.String()),
		zap.Duration("time", latency))

	status := c.Writer.Status()
	bodySize := c.Writer.Size()
	log.Info("Request processed",
		zap.String("Method", c.Request.Method),
		zap.String("URI", c.Request.URL.String()),
		zap.Int("Status", status),
		zap.Int("Body", bodySize))
}
