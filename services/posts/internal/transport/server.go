package transport

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/sulaman9009/microservices-sg/services/posts/internal/domain"
	"github.com/sulaman9009/microservices-sg/shared/logger/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":4000")
)

type server struct {
	mux        *echo.Echo
	logger     *zerolog.Logger
	post_store []*domain.Post
}

func NewHTTPServer(logger *zerolog.Logger) *server {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	return &server{
		mux:        e,
		logger:     logger,
		post_store: []*domain.Post{},
	}
}

func (s *server) Run() {
	s.mount()
	srv := &http.Server{
		Handler: s.mux,
		Addr:    httpAddr,
	}

	s.logger.Info().Msgf("Starting HTTP server on %s", httpAddr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// wait for shutdown signal
	if err := s.waitForShutdown(srv); err != nil {
		s.logger.Fatal().Err(err).Msgf("failed to shutdown server correctly")
	}
}

func (s *server) waitForShutdown(server *http.Server) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	// block here
	<-sig
	s.logger.Info().Msg("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	return nil
}
