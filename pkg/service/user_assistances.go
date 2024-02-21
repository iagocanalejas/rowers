package service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	a "rowers/templates/views/assistances"
	u "rowers/templates/views/users"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetUserAssistanceById(c echo.Context) error {
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

	assistance, err := s.db.GetUserAssistanceById(userID, assistanceID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Add("HX-Redirect", fmt.Sprintf("/users/%d/assistances/%d", userID, assistanceID))
	return a.UserAssistanceDetails(*assistance).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) GetUserAssistance(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	assistances, err := s.db.GetUserAssistancesByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserAssistancesTable(userID, assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) AddUserAssistance(c echo.Context) error {
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

	assistances, err := s.db.GetUserAssistancesByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserAssistancesTable(userID, assistances).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) DeleteUserAssistance(c echo.Context) error {
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

	assistances, err := s.db.GetUserAssistancesByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserAssistancesTable(userID, assistances).Render(c.Request().Context(), c.Response().Writer)
}
