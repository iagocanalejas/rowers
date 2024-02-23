package service

import (
	"fmt"
	"net/http"

	a "rowers/templates/views/assistances"
	u "rowers/templates/views/users"

	"github.com/labstack/echo/v4"
)

type UserAndAssistanceParams struct {
	UserID       int64 `query:"user_id" param:"user_id" json:"user_id"`
	AssistanceID int64 `query:"assistance_id" param:"assistance_id" json:"assistance_id"`
}

func (s *Service) GetUserAssistanceByID(c echo.Context) error {
	params := new(UserAndAssistanceParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistance, err := s.db.GetUserAssistanceByUserIDAndAssistanceID(params.UserID, params.AssistanceID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/users/%d/assistances/%d", params.UserID, params.AssistanceID))
	return a.UserAssistanceDetails(*assistance).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) GetUserAssistance(c echo.Context) error {
	params := new(UserParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetUserAssistancesByUserID(params.UserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserAssistancesTable(params.UserID, assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) CreateUserAssistance(c echo.Context) error {
	body := new(UserAndAssistanceParams)
	if err := c.Bind(body); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Info("adding assistance to user %d", body.UserID)
	if _, err := s.db.CreateUserAssistance(body.UserID, body.AssistanceID); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetUserAssistancesByUserID(body.UserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserAssistancesTable(body.UserID, assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) DeleteUserAssistance(c echo.Context) error {
	params := new(UserAndAssistanceParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Info("removing assistance %d from user %d", params.AssistanceID, params.UserID)
	if err := s.db.DeleteUserAssistance(params.UserID, params.AssistanceID); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetUserAssistancesByUserID(params.UserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserAssistancesTable(params.UserID, assistances).Render(c.Request().Context(), c.Response().Writer)
}
