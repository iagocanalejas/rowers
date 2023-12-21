package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) getUsers(c echo.Context) error {
	users, err := s.db.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
