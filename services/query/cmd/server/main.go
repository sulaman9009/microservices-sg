package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/sulaman9009/microservices-sg/services/query/internal/domain"
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

	query_store := map[string]*domain.Post{}

	e.GET("/posts", func(c echo.Context) error {
		posts := []*domain.Post{}
		for _, v := range query_store {
			posts = append(posts, v)
		}
		return c.JSON(http.StatusOK, domain.GetPostsResponse{
			Posts: posts,
		})
	})

	e.POST("/events", func(c echo.Context) error {
		var ev events.Event
		if err := c.Bind(&ev); err != nil {
			return err
		}
		fmt.Printf("query svc received event %s\n", ev.Type)
		switch ev.Type {
		case events.TypeCommentCreated:
			var data domain.CommentCreatedEventData
			if err := json.Unmarshal(ev.Data, &data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid comment created event")
			}
			if post, exists := query_store[data.PostId]; exists {
				post.Comments = append(post.Comments, &domain.Comment{
					Id:      data.Id,
					Content: data.Content,
					Status:  data.Status,
				})
			} else {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("post %s does not exist", data.PostId))
			}
		case events.TypePostCreated:
			var data domain.PostCreatedEventData
			if err := json.Unmarshal(ev.Data, &data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid post created event")
			}
			query_store[data.Id] = &domain.Post{
				Id:       data.Id,
				Title:    data.Title,
				Comments: []*domain.Comment{},
			}
		case events.TypeCommentUpdated:
			var data domain.CommentUpdatedEvent
			if err := json.Unmarshal(ev.Data, &data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid comment updated event")
			}
			post, ok := query_store[data.PostId]
			if !ok {
				err := fmt.Sprintf("could not find post %s for updated comment", data.PostId)
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			for _, comment := range post.Comments {
				if comment.Id == data.Id {
					comment.Status = data.Status
				}
			}
		}
		return nil
	})

	port := "4002"
	logger.Info().Msgf("Starting HTTP server on %s", port)
	return e.Start(fmt.Sprintf(":%s", port))
}
