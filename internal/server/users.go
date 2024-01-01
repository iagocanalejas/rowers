package server

import (
	"log"
	"net/http"
	"strconv"

	"rowers/internal/database"
	"rowers/internal/views"

	"github.com/labstack/echo/v4"
)

// TODO: error page
func (s *Server) getUsers(c echo.Context) error {
	users, err := s.db.GetUsers()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return views.Users(users).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) createUser(c echo.Context) error {
	user := new(database.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println("creating new user")
	user, err := s.db.CreateUser(*user)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSON(http.StatusCreated, user)
	}
	return views.UserRow(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) deleteUser(c echo.Context) error {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("removing user %d", user_id)
	if err := s.db.DeleteUser(user_id); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
