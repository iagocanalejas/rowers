package server

import (
	"log"
	"net/http"
	"strconv"

	"rowers/internal/views"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetUserAssistance(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetAssistanceByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserAssistancesTable(userID, assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) AddUserAssistance(c echo.Context) error {
	assistanceData := new(struct {
		AssistanceID int64 `json:"assistance_id"`
	})
	if err := c.Bind(assistanceData); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println("adding assistance to user")
	if err := s.db.AddUserAssistance(userID, assistanceData.AssistanceID); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetAssistanceByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserAssistancesTable(userID, assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) DeleteUserAssistance(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	assistanceID, err := strconv.ParseInt(c.Param("assistance_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("removing assistance %d from user %d", assistanceID, userID)
	if err := s.db.DeleteUserAssistance(userID, assistanceID); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetAssistanceByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserAssistancesTable(userID, assistances).Render(c.Request().Context(), c.Response().Writer)
}
