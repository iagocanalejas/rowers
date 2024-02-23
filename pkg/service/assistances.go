package service

import (
	"database/sql"
	"net/http"
	"time"

	"rowers/internal/db"
	d "rowers/templates/views/dashboard"

	"github.com/labstack/echo/v4"
)

type AssistanceParams struct {
	AssistanceID int64 `query:"assistance_id" param:"assistance_id" json:"assistance_id"`
}

type CreateAssistanceBody struct {
	Date string `json:"date"`
	Type string `json:"type"`
}

func (s *Service) GetAssistances(c echo.Context) error {
	assistances, err := s.db.GetAssistances()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return d.AssistancesTable(assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) CreateAssistance(c echo.Context) error {
	body := new(CreateAssistanceBody)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	parsedTime, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date format")
	}

	c.Logger().Info("creating new assistance")
	assistance, err := s.db.CreateAssistance(db.Assistance{Date: sql.NullTime{Time: parsedTime}, Type: body.Type})
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return d.AssistanceRow(*assistance).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) DeleteAssistance(c echo.Context) error {
	params := new(AssistanceParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Info("removing assistance %d", params.AssistanceID)
	if err := s.db.DeleteAssistance(params.AssistanceID); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
