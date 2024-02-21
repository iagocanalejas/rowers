package service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"rowers/internal/db"
	u "rowers/templates/views/users"
	d "rowers/templates/views/dashboard"

	"github.com/labstack/echo/v4"
)

// TODO: error page
func (s *Service) GetUserById(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.db.GetUserById(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/users/%d", userID))
	return u.UserDetails(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) GetUsers(c echo.Context) error {
	users, err := s.db.GetUsers()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return d.UsersTable(users).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) CreateUser(c echo.Context) error {
	user := new(db.User)
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
	return d.UserRow(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) DeleteUser(c echo.Context) error {
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
