package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/odrianoaliveira/go-service/internal/database"
	"github.com/odrianoaliveira/go-service/models"
	"log"
	"net/http"
)

type (
	Server interface {
		Start() error
		Readiness(ctx echo.Context) error
		Liveness(ctx echo.Context) error
	}
)

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()

	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server shutdown ocurred: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		health := models.Health{Status: "OK"}
		return ctx.JSON(http.StatusOK, health)
	}
	return ctx.JSON(http.StatusServiceUnavailable, models.Health{Status: "Unavailable"})
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/liveness", s.Liveness)
	s.echo.GET("/readiness", s.Readiness)
}
