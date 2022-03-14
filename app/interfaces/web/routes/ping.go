package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type pingHandler struct {
}

func NewPingHandler(e *echo.Echo) {
	h := &pingHandler{}
	e.GET("/ping", h.Ping)
}

func (p *pingHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
