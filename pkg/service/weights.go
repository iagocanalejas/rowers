package service

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"rowers/internal/db"
	u "rowers/templates/views/users"

	"github.com/labstack/echo/v4"
)

// TODO: add date-range
func (s *Service) GetUserWeights(c echo.Context) error {
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

	return u.UserWeights(userID, weights).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) AddWeight(c echo.Context) error {
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

	return u.UserWeights(weightData.UserID, weights).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) DeleteWeight(c echo.Context) error {
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

	return u.UserWeights(userID, weights).Render(c.Request().Context(), c.Response().Writer)
}

func toUserWeight(c echo.Context) (*db.Weight, error) {
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

	return &db.Weight{UserID: userId, Weight: weight}, nil
}
