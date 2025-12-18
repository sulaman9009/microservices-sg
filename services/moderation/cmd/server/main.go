package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/sulaman9009/microservices-sg/services/moderation/internal/domain"
	"github.com/sulaman9009/microservices-sg/shared/events"
	"github.com/sulaman9009/microservices-sg/shared/logger"
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
		var ev events.Event
		if err := c.Bind(&ev); err != nil {
			return err
		}
		fmt.Printf("moderation svc received event %s\n", ev.Type)
		switch ev.Type {
		case events.TypeCommentCreated:
			var data domain.CommentCreatedEventData
			if err := json.Unmarshal(ev.Data, &data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid comment created event")
			}
			var status string
			if strings.Contains(data.Content, "orange") {
				status = "rejected"
			} else {
				status = "approved"
			}
			data.Status = status
			go func() {
				time.Sleep(5 * time.Second)
				if err := events.EmitEvent(events.TypeCommentModerated, data); err != nil {
					fmt.Println("error emitting moderation event:", err)
				}
			}()
		}
		return nil
	})

	port := "4003"
	logger.Info().Msgf("Starting HTTP server on %s", port)
	return e.Start(fmt.Sprintf(":%s", port))
}
