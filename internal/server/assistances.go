package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"rowers/internal/database"
	"rowers/internal/views"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetAssistances(c echo.Context) error {
	assistances, err := s.db.GetAssistances()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return views.AssistancesTable(assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) CreateAssistance(c echo.Context) error {
	body := new(struct {
		Date string `json:"date"`
		Type string `json:"type"`
	})
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	parsedTime, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date format")
	}

	log.Println("creating new assistance")
	assistance, err := s.db.CreateAssistance(database.Assistance{Date: &parsedTime, Type: body.Type})
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSON(http.StatusCreated, assistance)
	}
	return views.AssistanceRow(*assistance).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) DeleteAssistance(c echo.Context) error {
	assistanceID, err := strconv.ParseInt(c.Param("assistance_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("removing assistance %d", assistanceID)
	if err := s.db.DeleteAssistance(assistanceID); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
