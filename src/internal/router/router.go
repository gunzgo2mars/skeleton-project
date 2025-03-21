package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *HttpServer) initRouting() {
	root := s.server

	root.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok!")
	})

	v1 := root.Group("/api/v1")

	v1.POST("/create-user", nil)
	v1.GET("/user", nil)
}
