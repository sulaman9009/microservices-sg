package main

import (
	"github.com/rs/zerolog"
	"github.com/sulaman9009/microservices-sg/services/posts/internal/transport"
	"github.com/sulaman9009/microservices-sg/shared/logger"
)

func main() {
	logger := logger.New()
	if err := run(logger); err != nil {
		logger.Fatal().Err(err)
	}
}

func run(logger *zerolog.Logger) error {
	srv := transport.NewHTTPServer(logger)
	srv.Run()
	return nil
}
