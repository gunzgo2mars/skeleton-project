package router

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	server *echo.Echo
	port   string
}

func New(port string) *HttpServer {
	e := echo.New()

	httpServer := &HttpServer{
		server: e,
		port:   port,
	}

	httpServer.initRouting()

	return httpServer
}

func (s *HttpServer) Start() error {
	return s.server.Start(fmt.Sprintf(":%s", s.port))
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *HttpServer) Server() *echo.Echo {
	return s.server
}
