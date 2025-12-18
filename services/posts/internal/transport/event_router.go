package transport

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sulaman9009/microservices-sg/shared/events"
)

func (s *server) eventHandler(c echo.Context) error {
	var ev events.Event
	if err := c.Bind(&ev); err != nil {
		return err
	}
	fmt.Printf("post svc received event %s\n", ev.Type)
	return nil
}
