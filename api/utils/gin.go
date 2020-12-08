package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fanap-infra/log"
	"github.com/gin-gonic/gin"
)

// ScopeLogGin ...
const ScopeLogGin = "GIN"

// NewGinEngine instead gin.Default()
func NewGinEngine(releaseMode bool) *gin.Engine {
	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(ginLogger(), gin.Recovery())
	return engine
}

func ginLogger() gin.HandlerFunc {
	logger := log.GetCustom(ScopeLogGin, 1)

	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// statusColor := colorForStatus(statusCode)
		// methodColor := colorForMethod(method)
		path := c.Request.URL.Path

		switch {
		case statusCode >= 400 && statusCode <= 499:
			err := c.Errors.String()
			if err != "" {
				logger.Warnv("Request", "statusCode", statusCode, "latency", latency.String(), "clientIP", clientIP, "method", method, "path", path, "error", err)
			} else {
				logger.Warnv("Request", "statusCode", statusCode, "latency", latency.String(), "clientIP", clientIP, "method", method, "path", path)
			}
		case statusCode >= 500:
			err := c.Errors.String()
			if err != "" {
				logger.Errorv("Request", "statusCode", statusCode, "latency", latency.String(), "clientIP", clientIP, "method", method, "path", path, "error", err)
			} else {
				logger.Errorv("Request", "statusCode", statusCode, "latency", latency.String(), "clientIP", clientIP, "method", method, "path", path)
			}
		default:
			err := c.Errors.String()
			if err != "" {
				logger.Infov("Request", "statusCode", statusCode, "latency", latency.String(), "clientIP", clientIP, "method", method, "path", path, "error", err)
			} else {
				logger.Infov("Request", "statusCode", statusCode, "latency", latency.String(), "clientIP", clientIP, "method", method, "path", path)
			}
		}
	}
}

func RunGin(engine *gin.Engine, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorv("Error on listen HTTP Server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("") // skip ^C
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
