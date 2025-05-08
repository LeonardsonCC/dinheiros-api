package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/handler"
	"github.com/LeonardsonCC/dinheiros/internal/logger"
	"github.com/LeonardsonCC/dinheiros/internal/profiling"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}
}

func run() error {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := profiling.SetupPyroscope()
	if err != nil {
		return err
	}

	// Set up OpenTelemetry.
	otelShutdown, err := telemetry.SetupOTelSDK(ctx)
	if err != nil {
		return err
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	r := setupServer(ctx)

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- r.Run()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return err
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	return nil
}

func setupServer(ctx context.Context) *gin.Engine {
	// service
	_, err := db.GetConnection(ctx)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(otelgin.Middleware("dinheiros-api"))
	r.Use(logger.Middleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	for _, route := range handler.Routes {
		route(r)
	}

	return r
}
