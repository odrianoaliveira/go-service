package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("email")
	customers, err := s.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, customers)
}
