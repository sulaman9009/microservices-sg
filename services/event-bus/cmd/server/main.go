package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/sulaman9009/microservices-sg/shared/logger"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := logger.New()
	if err := run(logger); err != nil {
		logger.Fatal().Err(err).Msg("failed to start server")
	}
}

func run(logger *zerolog.Logger) error {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/events", func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "could not read request")
		}
		return callEventConsumers([]string{
			"http://localhost:4000/events",
			"http://localhost:4001/events",
			"http://localhost:4002/events",
		}, body)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "working")
	})

	port := "4005"
	logger.Info().Msgf("Starting HTTP server on %s", port)
	return e.Start(fmt.Sprintf(":%s", port))
}

func callEventConsumers(addr []string, body []byte) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	for _, address := range addr {
		address := address // capture range variable
		g.Go(func() error {
			return callEventConsumer(ctx, address, body)
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func callEventConsumer(ctx context.Context, address string, body []byte) error {
	// create a request that respects the provided context so it can be cancelled
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, address, bytes.NewReader(body))
	if err != nil {
		return err
	}
	// assume events are JSON
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d when calling %s: %s", resp.StatusCode, address, string(respBody))
	}
	return nil
}
