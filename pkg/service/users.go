package service

import (
	"fmt"
	"net/http"

	"rowers/internal/db"
	d "rowers/templates/views/dashboard"
	u "rowers/templates/views/users"

	"github.com/labstack/echo/v4"
)

type UserParams struct {
	UserID int64 `query:"user_id" param:"user_id" json:"user_id"`
}

// TODO: error page
func (s *Service) GetUserByID(c echo.Context) error {
	params := new(UserParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.db.GetUserByID(params.UserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/users/%d", params.UserID))
	return u.UserDetails(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) GetUsers(c echo.Context) error {
	users, err := s.db.GetUsers()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return d.UsersTable(users).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) CreateUser(c echo.Context) error {
	user := new(db.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Info("create new user")
	user, err := s.db.CreateUser(*user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return d.UserRow(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) DeleteUser(c echo.Context) error {
	params := new(UserParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Info("removing user %d", params.UserID)
	if err := s.db.DeleteUser(params.UserID); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
