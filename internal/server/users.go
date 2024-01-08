package server

import (
	"fmt"
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
	return views.UserTable(users).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) getUserById(c echo.Context) error {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.db.GetUserById(user_id)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/users/%d", user_id))
	return views.UserDetails(*user).Render(c.Request().Context(), c.Response().Writer)
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

	// HACK: needs to return the deleted HTML so HTMX can update the DOM
	return views.UserRow(database.User{Id: user_id}).Render(c.Request().Context(), c.Response().Writer)
}
