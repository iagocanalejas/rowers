package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"rowers/internal/database"
	"rowers/internal/views"

	"github.com/labstack/echo/v4"
)

// TODO: add date-range
func (s *Server) GetWeightsByUserId(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	weights, err := s.db.GetWeightsByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserWeights(userID, weights).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) AddWeight(c echo.Context) error {
	weightData, err := toUserWeight(c)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println("adding weight to user")
	if err := s.db.AddWeight(weightData.UserID, weightData.Weight); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	weights, err := s.db.GetWeightsByUserId(weightData.UserID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserWeights(weightData.UserID, weights).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) DeleteWeight(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	weightID, err := strconv.ParseInt(c.Param("weight_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("removing weight %d from user %d", weightID, userID)
	if err := s.db.DeleteWeight(userID, weightID); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	weights, err := s.db.GetWeightsByUserId(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserWeights(userID, weights).Render(c.Request().Context(), c.Response().Writer)
}

func toUserWeight(c echo.Context) (*database.Weight, error) {
	weightData := new(struct {
		Weight string `json:"weight"`
	})
	if err := c.Bind(weightData); err != nil {
		log.Println(err)
		return nil, err
	}

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	weight, err := strconv.ParseFloat(weightData.Weight, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if weight < 30 || weight > 200 {
		return nil, errors.New("invalid weight")
	}

	return &database.Weight{UserID: userId, Weight: weight}, nil
}
