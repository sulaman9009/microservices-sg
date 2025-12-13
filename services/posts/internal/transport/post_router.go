package transport

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sulaman9009/microservices-sg/services/posts/internal/domain"
)

func (s *server) getPosts(c echo.Context) error {
	resp := domain.GetPostsResponse{
		Posts: s.post_store,
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *server) createPost(c echo.Context) error {
	var req domain.CreatePost
	if err := c.Bind(&req); err != nil {
		return err
	}
	rand, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	post := &domain.Post{
		Id:    rand.String(),
		Title: req.Title,
	}
	s.post_store = append(s.post_store, post)
	return c.JSON(http.StatusCreated, post)
}
