package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/sulaman9009/microservices-sg/services/comments/internal/domain"
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

	comment_store := map[string][]*domain.Comment{}

	e.GET("/posts/:id/comments", func(c echo.Context) error {
		postId := c.Param("id")
		if comments, exist := comment_store[postId]; exist {
			resp := domain.GetCommentsResp{
				Comments: comments,
			}
			return c.JSON(http.StatusOK, resp)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "post does not exist")
		}
	})

	e.POST("/posts/:id/comment", func(c echo.Context) error {
		fmt.Println("called")
		var req domain.CreateCommentReq
		if err := c.Bind(&req); err != nil {
			return err
		}
		rand, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		newComment := domain.Comment{
			Id:      rand.String(),
			Content: req.Content,
		}
		postId := c.Param("id")
		if comments, exists := comment_store[postId]; exists {
			comment_store[postId] = append(comments, &newComment)
		} else {
			comment_store[postId] = []*domain.Comment{&newComment}
		}
		return c.JSON(http.StatusCreated, newComment)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "working")
	})

	port := "4001"
	logger.Info().Msgf("Starting HTTP server on %s", port)
	return e.Start(fmt.Sprintf(":%s", port))
}
