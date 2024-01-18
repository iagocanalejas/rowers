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
func (s *Server) GetUserByID(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.db.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/users/%d", userID))
	return views.UserDetails(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) GetUsers(c echo.Context) error {
	users, err := s.db.GetUsers()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return views.UsersTable(users).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) CreateUser(c echo.Context) error {
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

func (s *Server) DeleteUser(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("removing user %d", userID)
	if err := s.db.DeleteUser(userID); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
